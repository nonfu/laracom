package repo

import (
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/go-test/deep"
    "github.com/jinzhu/gorm"
    "github.com/nonfu/laracom/user-service/model"
    "github.com/stretchr/testify/require"
    "github.com/stretchr/testify/suite"
    "testing"
)

type MockUserRepoSuite struct {
    suite.Suite                  // 辅助测试
    Db *gorm.DB                  // 模拟gorm
    mock sqlmock.Sqlmock         // 保存测试接口
    repo *UserRepository          // Repository 实现
    model *model.User            // User 模型
}

func (s *MockUserRepoSuite) SetupSuite() {
    db, mock, err := sqlmock.New()
    s.mock = mock
    require.NoError(s.T(), err)
    // 使用 mock 的虚拟化 db 初始化 gorm 连接
    s.Db, err = gorm.Open("mysql", db)
    require.NoError(s.T(), err)
    s.Db.LogMode(true)
    s.repo = &UserRepository{Db: s.Db}
}

func TestInit(t *testing.T) {
    suite.Run(t, new(MockUserRepoSuite))
}

func (s *MockUserRepoSuite) AfterTest(_, _ string) {
    require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *MockUserRepoSuite) Test_repository_Create() {
    var (
        name = "test"
        email = "test@xueyuanjun.com"
        status = 1
    )

    s.mock.ExpectBegin()
    s.mock.ExpectExec("INSERT INTO `users`").WillReturnResult(sqlmock.NewResult(1, 1))
    s.mock.ExpectCommit()

    user := &model.User{Name: name, Email: email, Status: uint8(status)}
    err := s.repo.Create(user)

    require.NoError(s.T(), err)
}

func (s *MockUserRepoSuite) Test_repository_Get() {
    var id uint = 1

    s.mock.ExpectQuery("SELECT *").WithArgs(id).
        WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(id, "test", "test@xueyuanjun.com"))

    res, err := s.repo.Get(id)

    require.NoError(s.T(), err)
    require.Nil(s.T(), deep.Equal(id, res.ID))
}