function firebase_init() {
    var firebaseConfig = {
        apiKey: "AIzaSyDRMQ-tfWZ1Ip_5NnSE9ekMcC_GyYG-S_4",
        authDomain: "travelblog-4b7fa.firebaseapp.com",
        databaseURL: "https://travelblog-4b7fa.firebaseio.com",
        projectId: "travelblog-4b7fa",
        storageBucket: "travelblog-4b7fa.appspot.com",
        messagingSenderId: "468125028366",
        appId: "1:468125028366:web:9932d6df907559cfa52c5f",
        measurementId: "G-6KG8YQ476S"
    };
    firebase.initializeApp(firebaseConfig);
    firebase.analytics();
}
