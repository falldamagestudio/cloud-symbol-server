
# Set up infrastructure

## Decide on names and locations

You need to make some decisions early on:
* If you want to have this tied to a static DNS name, What DNS hostname would you like the services to be present at?
* Which location do you want the system to run in? Which GCP region + zone? Pick something close to you.

## Google Cloud setup

* Create a user and an organization in Google Cloud Platform.
* Set up a billing account for your organization. Provide credit card information.
* Create a new project in Google Cloud Platform (GCP). Name it as `<company>-cloud-symbol-server`.
* Visit APIs & Services | OAuth consent screen. Set up an Internal screen. Name it after the GCP project. Add your company's domain as an authorized domain. [TODO: Verify whether this is correct.] Choose no scopes.
* Create a bucket for Terraform state storage within the new project. Name it as `<GCP project>-state`. Place it in the same region that you intend to have other resources. Choose Standard default storage class. Choose "Enforce public access prevention on this bucket". Choose Uniform access control. Enable Object Versioning, without any numeric limitations.

* Configure `gcloud` to have a configuration with the same name as the project. Set the default parameters (see `gcloud config`) and log in.

## Firebase setup

* Visit [Firebase Console](https://console.firebase.google.com/). Create a Firebase for your newly-created GCP project. Enable Google Analytics if you wish - it's not necessary though.
* Add a Web app. Also set up Hosting for app. Register app, then abort setup process.
* Set up Authentication. Choose a provider - Google, if you want people to use company Google accounts - and enable that.

## GitHub setup

* Create a GitHub organization for your company.
* Decide on how you want to handle your environment and deployment automation; either fork this repository,
  putting your new environment into `environments/<your company>` in the fork, or create a new repository that contains just your company's environment config, and has a submodule reference to this repo.

# First-time system bring-up

* Duplicate `environments/test/` to a new environment folder named `environments/<your company>`.
* Modify the contents of `**/backend.tf`, `**/terraform.tfvars` and `firebase/contents.json`; in particular, any strings that state `test-cloud-symbol-server` or will need to be changed, and any strings containing `europe-west` as well.
* Extract the Project Config from Firebase Console and insert into `firebase/frontend/firebase-config.json`.

* Run `make ENV=<path to your new environment folder> deploy`.

* Inspect in GCP Console what the access URLs are for the two cloud functions. Update `config.json` accordingly.

# Static DNS setup

* Visit "Hosting" in Firebase Console. Choose to add a custom domain.
* Enter a suitable domain name, such as `symbols.<your organization domain>`.
* Configure DNS entries with your DNS registrar.
* Visit "Authentication". Add the domain as an authorized domain under your authorization provider, if necessary.
