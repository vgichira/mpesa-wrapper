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
)
