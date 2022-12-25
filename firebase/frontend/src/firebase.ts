import { initializeApp } from 'firebase/app'
import { getAuth, onAuthStateChanged } from "firebase/auth"
import { getFirestore } from "firebase/firestore"

import store from './store/index'

import { firebaseConfig } from './firebaseConfig'

export const firebaseApp = initializeApp(firebaseConfig)

// When the Firebase SDK initializes, it does initally not know whether it
//  has a user since the previous session (or whether this is the page reload
//  just after a sign-in/sign-out operation)
// Once the SDK completes initialization, it will trigger an onAuthStateChanged()
//  callback, and provide either a user object or null.
// We use the login state of LoginState.Unknown to signal to the rest of our application
//  that it should not yet show any application UI.

store.commit('setLoginStateUnknown')

export const auth = getAuth(firebaseApp)
onAuthStateChanged(auth, user => {
    if(user){
        store.commit('setUser', user);
    }else{
        store.commit('setUser', null);
    }
});

export const db = getFirestore(firebaseApp)
