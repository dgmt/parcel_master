deploy:
	# gcloud auth login
	gcloud config set project parcel-master-402212
	gcloud app deploy --quiet

test:
	go test ./... -v