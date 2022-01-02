
import firebase from 'firebase/app'
import 'firebase/firestore'

import store from './store/index'

import { firebaseConfig, authEmulatorUrl, firestoreEmulator } from './firebaseConfig'

firebase.initializeApp(firebaseConfig)

// When the Firebase SDK initializes, it does initally not know whether it
//  has a user since the previous session (or whether this is the page reload
//  just after a sign-in/sign-out operation)
// Once the SDK completes initialization, it will trigger an onAuthStateChanged()
//  callback, and provide either a user object or null.
// We use the login state of LoginState.Unknown to signal to the rest of our application
//  that it should not yet show any application UI.

store.commit('setLoginStateUnknown')

firebase.auth().onAuthStateChanged((user) =>{
    if(user){
        store.commit('setUser', user);
    }else{
        store.commit('setUser', null);
    }
});

export const db = firebase.firestore()

if (firestoreEmulator) {
    db.useEmulator(firestoreEmulator.host, firestoreEmulator.port)
}

export const auth = firebase.auth()

if (authEmulatorUrl) {
    auth.useEmulator(authEmulatorUrl)
}
