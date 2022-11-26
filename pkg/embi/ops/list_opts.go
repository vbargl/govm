package ops

func setHttpClientToListOp(value HttpClient) ListOpt {
	return func(op *ListOp) {
		op.httpClient = value
	}
}
