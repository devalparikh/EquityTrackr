package com.EquityTrackr.backend.investor;

import com.google.api.core.ApiFuture;
import com.google.cloud.firestore.DocumentReference;
import com.google.cloud.firestore.DocumentSnapshot;
import com.google.cloud.firestore.Firestore;
import com.google.cloud.firestore.WriteResult;
import com.google.firebase.cloud.FirestoreClient;
import org.springframework.stereotype.Service;

import java.util.concurrent.ExecutionException;

//CRUD operations
@Service
public class InvestorService {

    public static final String COL_NAME = "investors";

    public Investor saveInvestor(Investor investor) throws InterruptedException, ExecutionException {
        Firestore dbFirestore = FirestoreClient.getFirestore();
//        ApiFuture<WriteResult> collectionsApiFuture = dbFirestore.collection(COL_NAME).document(investor.getName()).set(investor);
//        return collectionsApiFuture.get().getUpdateTime().toString();

        if (investor.getBalance() == null) {
            investor.setBalance((long) 0);
        }

        ApiFuture<WriteResult> collectionsApiFuture =
                dbFirestore
                        .collection(COL_NAME)
                        .document(investor.getName())
                        .set(investor);

        return investor;
    }

    public Investor getInvestor(String name) throws InterruptedException, ExecutionException {
        Firestore dbFirestore = FirestoreClient.getFirestore();
        DocumentReference documentReference = dbFirestore.collection(COL_NAME).document(name);
        ApiFuture<DocumentSnapshot> future = documentReference.get();

        DocumentSnapshot document = future.get();

        Investor investor = null;

        if (document.exists()) {
            investor = document.toObject(Investor.class);
            return investor;
        } else {
            return null;
        }
    }

    public String updateInvestor(Investor person) throws InterruptedException, ExecutionException {
        Firestore dbFirestore = FirestoreClient.getFirestore();
        ApiFuture<WriteResult> collectionsApiFuture = dbFirestore.collection(COL_NAME).document(person.getName()).set(person);
        return collectionsApiFuture.get().getUpdateTime().toString();
    }

    public String deleteInvestor(String name) {
        Firestore dbFirestore = FirestoreClient.getFirestore();
        ApiFuture<WriteResult> writeResult = dbFirestore.collection(COL_NAME).document(name).delete();
        return "Document with Investor ID " + name + " has been deleted";
    }

}