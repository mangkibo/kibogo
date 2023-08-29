package libraries

type (
	Messaging struct {
		FirebaseToken string
		Message       string
		Title         string
	}
)

func (m *Messaging) send() {
	//firebaseKey := Env("FIREBASE_KEY")

	//ctx := context.Background()
	//client, err := app.Messaging(ctx)
}
