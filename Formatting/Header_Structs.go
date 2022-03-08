package Formatting

type Master_Json struct {
	Sender_recipient_subject Sender_Recipient_Subject_Json
	Authentication           Authentictation_Json
	Additional_headers       Additional_Information_Json
}

type Authentictation_Json struct {
	X_SONIC_DKIM_SIGN          string   `json:"sonicDKIMsig,omitempty"`
	Authentication_Results     string   `json:"authresults,omitempty"`
	ARC_Authentication_Results string   `json:"arcauthresults,omitempty"`
	ARC_Seal                   string   `json:"arcseal,omitempty"`
	ARC_Message_Signature      string   `json:"arc_mes_sig,omitempty"`
	X_Google_DKIM_Signature    string   `json:"google_dkim_sig,omitempty"`
	Received_Path              []string `json:"receivedPath,omitempty"`
	Received_SPF               []string `json:"sonicDKIMsig,omitempty"`
	DKIM_Signature             []string `json:"sonicDKIMsig,omitempty"`
	Analysis_Score             int64    `json:"analysisScore"`
}

// add pgp encryption
type Sender_Recipient_Subject_Json struct {
	Sender_name    string `json:"SenderName,omitempty"`
	Sender_email   string `json:"SenderEmail,omitempty"`
	Rec_name       string `json:"RecipientName,omitempty"`
	Rec_email      string `json:"RecipientEmail,omitempty"`
	Subject        string `json:"Subject,omitempty"`
	Date           string `json:"SendDate,omitempty"`
	Analysis_Score int64  `json:"analysisScore"`
}

// lots to add to this, may also want to break into different groups in addition.
type Additional_Information_Json struct {
	Mime_Type                 string `json:"MimeType,omitempty"`
	Message_Id                string `json:"mesId,omitempty"`
	X_Feedback_ID             string `json:"FeedbackId,omitempty"`
	X_Google_Smtp_Source      string `json:"GoogleSmtpSource,omitempty"`
	MIME_Version              string `json:"MimeVersion,omitempty"`
	Content_Disposition       string `json:"ContentDisposition,omitempty"`
	Content_Transfer_Encoding string `json:"ContentTransferEncoding,omitempty"`
	X_Mailer                  string `json:"XMailer,omitempty"`
	X_Gm_Message_State        string `json:"GMMessageState,omitempty"`
	Content_Type              string `json:"ContentType,omitempty"`
	Content_Length            string `json:"ContentLength,omitempty"`
}
