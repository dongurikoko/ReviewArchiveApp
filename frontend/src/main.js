// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
//import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyCUuB03U7o2ZH1ObbcTZ9w0fqvBRkCQAtE",
  authDomain: "review-app-711a5.firebaseapp.com",
  projectId: "review-app-711a5",
  storageBucket: "review-app-711a5.appspot.com",
  messagingSenderId: "711734477056",
  appId: "1:711734477056:web:da88d1b05697158deb33e5",
  measurementId: "G-6KY6F1GQ14"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
