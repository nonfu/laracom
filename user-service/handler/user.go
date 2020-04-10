package handler

import (
    "encoding/json"
    "errors"
    "github.com/jinzhu/gorm"
    "github.com/micro/go-micro/broker"
    "github.com/nonfu/laracom/user-service/model"
    pb "github.com/nonfu/laracom/user-service/proto/user"
    "github.com/nonfu/laracom/user-service/repo"
    "github.com/nonfu/laracom/user-service/service"
    "golang.org/x/crypto/bcrypt"
    "golang.org/x/net/context"
    "log"
    "strconv"
)

const topic  = "password.reset"

type UserService struct {
    Repo repo.Repository
    ResetRepo repo.PasswordResetInterface
    Token service.Authable
    PubSub broker.Broker
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
    var userModel *model.User
    var err error
    if req.Id != "" {
        id, _ := strconv.ParseUint(req.Id, 10, 64)
        userModel, err = srv.Repo.Get(uint(id))
    } else if req.Email != "" {
        userModel, err = srv.Repo.GetByEmail(req.Email)
    }
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    if userModel != nil {
        res.User, _ = userModel.ToProtobuf()
    }
    return nil
}

func (srv *UserService) GetById(ctx context.Context, req *pb.User, res *pb.Response) error {
    name := "test" + req.Id
    res.User = &pb.User{Id: req.Id, Name: name}
    return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
    users, err := srv.Repo.GetAll()
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    userItems := make([]*pb.User, len(users))
    for index, user := range users {
        userItem, _ := user.ToProtobuf()
        userItems[index] = userItem
    }
    res.Users = userItems
    return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
    // 对密码进行哈希加密
    hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    req.Password = string(hashedPass)
    userModel := &model.User{}
    user, _ := userModel.ToORM(req)
    if err := srv.Repo.Create(user); err != nil {
        return err
    }
    res.User, _ = user.ToProtobuf()
    return nil
}

func (srv *UserService) Update(ctx context.Context, req *pb.User, res *pb.Response) error {
    if req.Password != "" {
        // 如果密码字段不为空的话对密码进行哈希加密
        hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        req.Password = string(hashedPass)
    }
    if req.Id == "" {
        return errors.New("用户 ID 不能为空")
    }
    id, _ := strconv.ParseUint(req.Id, 10, 64)
    user, _ := srv.Repo.Get(uint(id))
    if err := srv.Repo.Update(user); err != nil {
        return err
    }
    res.User, _ = user.ToProtobuf()
    return nil
}

func (srv *UserService) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
    // 获取用户信息
    user, err := srv.Repo.GetByEmail(req.Email)
    if err != nil {
        return err
    }

    // 校验用户输入密码是否于数据库存储密码匹配
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return err
    }

    // 生成 jwt token
    token, err := srv.Token.Encode(user)
    if err != nil {
        return err
    }
    res.Token = token
    return nil
}

func (srv *UserService) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

    // 校验用户亲求中的token信息是否有效
    claims, err := srv.Token.Decode(req.Token)

    if err != nil {
        return err
    }

    if claims.User.ID == 0 {
        return errors.New("无效的用户")
    }

    res.Valid = true

    return nil
}

func (srv *UserService) CreatePasswordReset(ctx context.Context, req *pb.PasswordReset, res *pb.PasswordResetResponse) error {
    if req.Email == "" {
        return errors.New("邮箱不能为空")
    }
    resetModel := new(model.PasswordReset)
    passwordReset, _ := resetModel.ToORM(req)
    if err := srv.ResetRepo.Create(passwordReset); err != nil {
        return err
    }
    if passwordReset != nil {
        req, _ = passwordReset.ToProtobuf()
        if err := srv.publishEvent(req); err != nil {
            return err
        }
        res.PasswordReset = req
    }
    return nil
}

func (srv *UserService) DeletePasswordReset(ctx context.Context, req *pb.PasswordReset, res *pb.PasswordResetResponse) error {
    if req.Email == "" {
        return errors.New("邮箱不能为空")
    }
    reset, err := srv.ResetRepo.GetByEmail(req.Email)
    if err != nil {
        return errors.New("数据库查询出错")
    }
    if err := srv.ResetRepo.Delete(reset); err != nil {
        return err
    }
    res.PasswordReset = nil
    return nil
}

func (srv *UserService) ValidatePasswordResetToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
    // 校验用户亲求中的token信息是否有效
    if req.Token == "" {
        return errors.New("Token信息不能为空")
    }

    _, err := srv.ResetRepo.GetByToken(req.Token)
    if err != nil && err != gorm.ErrRecordNotFound {
        return errors.New("数据库查询异常")
    }

    if err == gorm.ErrRecordNotFound {
        res.Valid = false
    } else {
        res.Valid = true
    }
    return nil
}

func (srv *UserService) publishEvent(reset *pb.PasswordReset) error {
    // Marshal to JSON strings
    body, err := json.Marshal(reset)
    if err != nil {
        return err
    }
    // Create a broker message
    msg := &broker.Message{
        Header: map[string]string{
            "email": reset.Email,
        },
        Body: body,
    }
    // Publish message to broker
    if err := srv.PubSub.Publish(topic, msg); err != nil {
        log.Printf("[pub] failed: %v", err)
    }
    return nil
}