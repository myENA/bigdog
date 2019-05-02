package vsz

// API Version: v8_0

type WSGUploadStatisticstoFTPService struct {
    client *Client
}

func NewWSGUploadStatisticstoFTPService (client *Client) *WSGUploadStatisticstoFTPService {
    s := new(WSGUploadStatisticstoFTPService)
    s.client = client
    return s
}

func (ss *WSGService) WSGUploadStatisticstoFTPService () *WSGUploadStatisticstoFTPService {
    serv := new(WSGUploadStatisticstoFTPService)
    serv.client = ss.client
    return serv
}

