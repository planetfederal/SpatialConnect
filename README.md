
### Overview
SpatialConnect is a a collection of open-source libraries that enables developers to build mobile applications with geospatial capabilities.  This includes native SDKs for Android and iOS, Javascript bridge libraries for developing apps with React Native, other supporting libraries, and a reference implementation of "mobile a backend as a service" that works with the SDKs.  This repo is a landing page for developers who want to develop on the SpatialConnect project.

### Prerequisites

If you are going to develop apps for iOS, you need to run macOS, therefore the following instructions assume you are using a Mac.

First install the useful package manager for macOS, [Homebrew](https://brew.sh/)
```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)
```

Then install a few required prerequisite packages if they are not already installed.

```
brew install node
brew install git
brew install watchman
```

For Android development, use the installers to install [Java](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html),
and [Android Studio](https://developer.android.com/studio/install.html).

For iOS, download and install
[Xcode](https://itunes.apple.com/us/app/xcode/id497799835?mt=12) from the App Store.

Many of the projects also use Docker containers for running services locally,
so make sure you [install Docker](https://docs.docker.com/docker-for-mac/install/#download-docker-for-mac).

### SpatialConnect git repositories

There are several repos that you will need to clone before starting to develop for SpatialConnect.


**spatialconnect-android-sdk** - the native SDK for the Android platform.
```
git clone https://github.com/boundlessgeo/spatialconnect-android-sdk
```
After cloning this repo, you can open it in Android Studio and follow the prompts
to install the required Android SDK platforms and build tools.


**spatialconnect-ios-sdk** - the native SDK for the iOS platform.
```
git clone https://github.com/boundlessgeo/spatialconnect-ios-sdk
```
The iOS SDK uses [Carthage](https://github.com/Carthage/Carthage) to manage its
dependencies so install it with `brew install carthage` if you haven't already.


**spatialconnect-js** - the Javascript bridge which allows a Javascript
runtime (web view, React Native app, etc) to communicate to the native SDKs.
```
git clone https://github.com/boundlessgeo/spatialconnect-js
```


**spatialconnect-mobile** - a reference implementation of a React Native app that uses the native SDKs.
```
git clone https://github.com/boundlessgeo/spatialconnect-mobile
```
Make sure you installed the react-native-cli `npm install -g react-native-cli` so
you can use it to start the apps from the command line.

**spatialconnect-server** - a reference implementation of a mobile backend as a service
that includes user management, push notifications, device configuration, a form
builder for data collection, and other features.
```
git clone https://github.com/boundlessgeo/spatialconnect-server
```

There are a few other supporting repos that you might also need to clone,
depending on what you're developing.

**react-native-spatialconnect** - a Javascript library used to integrate SpatialConnect with  React Native applications
```
git clone https://github.com/boundlessgeo/react-native-spatialconnect
```

**spatialconnect-form-schema** - a Javascript library used to integrate with [tcomb-form](https://github.com/gcanti/tcomb-form) which allows us to render
forms based on JSON schemas.
```
git clone https://github.com/boundlessgeo/spatialconnect-form-schema
```

**schema** - a repo holding protobuf message schemas
```
git clone https://github.com/boundlessgeo/schema
```
