package mpesa

type (
	Config struct {
		ConsumerKey    string
		ConsumerSecret string
		Environment    string
	}

	LipaNaMpesaRequest struct {
		BusinessShortCode string `json:"BusinessShortCode"`
		Password          string `json:"Password"`
		Timestamp         string `json:"Timestamp"`
		TransactionType   string `json:"TransactionType"`
		Amount            string `json:"Amount"`
		PartyA            string `json:"PartyA"`
		PartyB            string `json:"PartyB"`
		PhoneNumber       string `json:"PhoneNumber"`
		CallBackURL       string `json:"CallBackURL"`
		AccountReference  string `json:"AccountReference"`
		TransactionDesc   string `json:"TransactionDesc"`
	}

	RegisterURL struct {
		ValidationURL   string `json:"ValidationURL"`
		ConfirmationURL string `json:"ConfirmationURL"`
		ResponseType    string `json:"ResponseType"`
		ShortCode       string `json:"ShortCode"`
	}

	C2BTransaction struct {
		CommandID     string `json:"CommandID"`
		Amount        string `json:"Amount"`
		MSISDN        string `json:"Msisdn"`
		BillRefNumber string `json:"BillRefNumber"`
		ShortCode     string `json:"ShortCode"`
	}

	B2C struct {
		InitiatorName      string `json:"InitiatorName"`
		SecurityCredential string `json:"SecurityCredential"`
		CommandID          string `json:"CommandID"`
		Amount             string `json:"Amount"`
		PartyA             string `json:"PartyA"`
		PartyB             string `json:"PartyB"`
		Remarks            string `json:"Remarks"`
		QueueTimeoutURL    string `json:"QueueTimeOutURL"`
		ResultURL          string `json:"ResultURL"`
		Occassion          string `json:"Occassion"`
	}

	Reversal struct {
		Initiator              string `json:"Initiator"`
		SecurityCredential     string `json:"SecurityCredential"`
		CommandID              string `json:"CommandID"`
		TransactionID          string `json:"TransactionID"`
		Amount                 string `json:"Amount"`
		ReceiverParty          string `json:"ReceiverParty"`
		ReceiverIdentifierType string `json:"ReceiverIdentifierType"`
		ResultURL              string `json:"ResultURL"`
		QueueTimeoutURL        string `json:"QueueTimeoutURL"`
		Remarks                string `json:"Remarks"`
		Occasion               string `json:"Occasion"`
	}

	TransactionStatus struct {
		Initiator          string `json:"Initiator"`
		SecurityCredential string `json:"SecurityCredential"`
		CommandID          string `json:"CommandID"`
		TransactionID      string `json:"TransactionID"`
		PartyA             string `json:"PartyA"`
		IdentifierType     string `json:"IdentifierType"`
		ResultURL          string `json:"ResultURL"`
		QueueTimeoutURL    string `json:"QueueTimeOutURL"`
		Remarks            string `json:"Remarks"`
		Occasion           string `json:"Occasion"`
	}
)
