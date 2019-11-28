package handler

import (
    "errors"
    "github.com/jinzhu/gorm"
    pb "github.com/nonfu/laracom/user-service/proto/user"
    "github.com/nonfu/laracom/user-service/repo"
    "github.com/nonfu/laracom/user-service/service"
    "golang.org/x/crypto/bcrypt"
    "golang.org/x/net/context"
)

type UserService struct {
    Repo repo.Repository
    ResetRepo repo.PasswordResetInterface
    Token service.Authable
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
    var user *pb.User
    var err error
    if req.Id != "" {
        user, err = srv.Repo.Get(req.Id)
    } else if req.Email != "" {
        user, err = srv.Repo.GetByEmail(req.Email)
    }
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    res.User = user
    return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
    users, err := srv.Repo.GetAll()
    if err != nil && err != gorm.ErrRecordNotFound {
        return err
    }
    res.Users = users
    return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
    // 对密码进行哈希加密
    hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    req.Password = string(hashedPass)
    if err := srv.Repo.Create(req); err != nil {
        return err
    }
    res.User = req
    return nil
}

func (srv *UserService) Update(ctx context.Context, req *pb.User, res *pb.Response) error {
    if req.Id == "" {
        return errors.New("用户 ID 不能为空")
    }
    if req.Password != "" {
        // 如果密码字段不为空的话对密码进行哈希加密
        hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        req.Password = string(hashedPass)
    }
    if err := srv.Repo.Update(req); err != nil {
        return err
    }
    res.User = req
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

    if claims.User.Id == "" {
        return errors.New("无效的用户")
    }

    res.Valid = true

    return nil
}

func (srv *UserService) CreatePasswordReset(ctx context.Context, req *pb.PasswordReset, res *pb.PasswordResetResponse) error {
    if req.Email == "" {
        return errors.New("邮箱不能为空")
    }
    if err := srv.ResetRepo.Create(req); err != nil {
        return err
    }
    res.PasswordReset = req
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