package repository

var globalRepo UserRepository

func InitGlobalRepo(r UserRepository) {
    globalRepo = r
}

func GetGlobalRepo() UserRepository {
    return globalRepo
}
