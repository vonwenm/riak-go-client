package riak

// Convert from csv with:
// %s/\(\d\+\),\([^,]\+\),.*/var rpbCode_\2 byte = \1/
var rpbCode_RpbErrorResp byte = 0
var rpbCode_RpbPingReq byte = 1
var rpbCode_RpbPingResp byte = 2
var rpbCode_RpbGetClientIdReq byte = 3
var rpbCode_RpbGetClientIdResp byte = 4
var rpbCode_RpbSetClientIdReq byte = 5
var rpbCode_RpbSetClientIdResp byte = 6
var rpbCode_RpbGetServerInfoReq byte = 7
var rpbCode_RpbGetServerInfoResp byte = 8
var rpbCode_RpbGetReq byte = 9
var rpbCode_RpbGetResp byte = 10
var rpbCode_RpbPutReq byte = 11
var rpbCode_RpbPutResp byte = 12
var rpbCode_RpbDelReq byte = 13
var rpbCode_RpbDelResp byte = 14
var rpbCode_RpbListBucketsReq byte = 15
var rpbCode_RpbListBucketsResp byte = 16
var rpbCode_RpbListKeysReq byte = 17
var rpbCode_RpbListKeysResp byte = 18
var rpbCode_RpbGetBucketReq byte = 19
var rpbCode_RpbGetBucketResp byte = 20
var rpbCode_RpbSetBucketReq byte = 21
var rpbCode_RpbSetBucketResp byte = 22
var rpbCode_RpbMapRedReq byte = 23
var rpbCode_RpbMapRedResp byte = 24
var rpbCode_RpbIndexReq byte = 25
var rpbCode_RpbIndexResp byte = 26
var rpbCode_RpbSearchQueryReq byte = 27
var rpbCode_RpbSearchQueryResp byte = 28
var rpbCode_RpbResetBucketReq byte = 29
var rpbCode_RpbResetBucketResp byte = 30
var rpbCode_RpbGetBucketTypeReq byte = 31
var rpbCode_RpbSetBucketTypeReq byte = 32
var rpbCode_RpbGetBucketKeyPreflistReq byte = 33
var rpbCode_RpbGetBucketKeyPreflistResp byte = 34
var rpbCode_RpbCSBucketReq byte = 40
var rpbCode_RpbCSBucketResp byte = 41
var rpbCode_RpbCounterUpdateReq byte = 50
var rpbCode_RpbCounterUpdateResp byte = 51
var rpbCode_RpbCounterGetReq byte = 52
var rpbCode_RpbCounterGetResp byte = 53
var rpbCode_RpbYokozunaIndexGetReq byte = 54
var rpbCode_RpbYokozunaIndexGetResp byte = 55
var rpbCode_RpbYokozunaIndexPutReq byte = 56
var rpbCode_RpbYokozunaIndexDeleteReq byte = 57
var rpbCode_RpbYokozunaSchemaGetReq byte = 58
var rpbCode_RpbYokozunaSchemaGetResp byte = 59
var rpbCode_RpbYokozunaSchemaPutReq byte = 60
var rpbCode_DtFetchReq byte = 80
var rpbCode_DtFetchResp byte = 81
var rpbCode_DtUpdateReq byte = 82
var rpbCode_DtUpdateResp byte = 83
var rpbCode_RpbAuthReq byte = 253
var rpbCode_RpbAuthResp byte = 254
var rpbCode_RpbStartTls byte = 255
