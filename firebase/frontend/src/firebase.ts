import { initializeApp } from 'firebase/app'
import { getAuth, onAuthStateChanged } from "firebase/auth"
import { getFirestore } from "firebase/firestore"

import { useAuthUserStore } from './stores/authUser'
import { firebaseConfig } from './appConfig'

export const firebaseApp = initializeApp(firebaseConfig)

export const auth = getAuth(firebaseApp)
onAuthStateChanged(auth, user => {
    const authUserStore = useAuthUserStore()
    if(user){
        authUserStore.setUser(user);
    }else{
        authUserStore.setUser(null);
    }
});

export const db = getFirestore(firebaseApp)
