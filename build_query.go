package gopg_mock

type buildQuery struct {
	funcName string
	query    string
	params   []interface{}
	result   *OrmResult
	err      error
}
