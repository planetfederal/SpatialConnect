
#### Fresh install of EFC and the associated dependencies.

#### Dependencies List
What you will need to have installed prior to running the program:

##### Git repositories
* [spatialconnect-mobile](https://github.com/boundlessgeo/spatialconnect-mobile)
`
git clone https://github.com/boundlessgeo/spatialconnect-mobile
`

##### Signing apps

First you will need to obtain the signing key file and passwords. Requests can be made via email at spatialconnect@boundlessgeo.com.


##### General


* Homebrew:
`
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)
`

* Node:
`
brew install node
`

* Watchman:
`
brew install watchman
`

* React-native CLI:
`
npm install -g react-native-cli
`

###### iOS specific
* Xcode
* Carthage
* Fabric
* Crashlytics.framework

###### Android specific
* Android Studio
* Java Development Kit 1.8
* gradle.properties

#### iOS setup
* Install Xcode from the Apple App store.

* Install Carthage
`
npm install carthage
`

* Make sure all three git repositories are on a common level in your directory structure.
* Add the Crashlytics.framework file into the _spatialconnect-mobile_ > _ios_ folder.
* In the _spatialconnect-mobile folder_
`
npm i
`

* Finally **either**
`
react-native run-ios
`

* **OR** in the _ios_ folder of _spatialconnect-mobile_ open the SCMobile.xcodeproj file.

#### Android setup
* Install [Android Studio](https://developer.android.com/studio/index.html)

* Next in the _spatialconnect-mobile folder_
`
npm i
`  

* To start the Android app, first you need to start an emulator.

* Finally:
`
react-native run-android
`
