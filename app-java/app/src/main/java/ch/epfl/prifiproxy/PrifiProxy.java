package ch.epfl.prifiproxy;

import android.app.Application;
import android.content.Context;
import android.content.SharedPreferences;

import servicego.Servicego;

/**
 * The entry point of this app
 * This class allows us to do some initialization work after launching the app.
 */
public class PrifiProxy extends Application {

    private static Application mApplication;

    @Override
    public void onCreate() {
        super.onCreate();
        mApplication = this;

        initPrifiConfig();
    }

    public static Context getContext() {
        return mApplication.getApplicationContext();
    }

    /**
     * PriFi Initialization
     * Retrieve, save and modify (if necessary) the important PriFi values
     */
    private void initPrifiConfig() {

        SharedPreferences preferences = getSharedPreferences(getString(R.string.prifi_config_shared_preferences), MODE_PRIVATE);
        Boolean isFirstInit = preferences.getBoolean(getString(R.string.prifi_config_first_init), true);

        // Save if it's the first initialization
        if (isFirstInit) {
            final String pubKey;
            final String priKey;

            try {
                Servicego.generateNewKeyPairAndAssign();
                pubKey = Servicego.getPublicKey();
                priKey = Servicego.getPrivateKey();
            } catch (Exception e) {
                throw new RuntimeException("Can't generate keys");
            }

            SharedPreferences.Editor editor = getSharedPreferences(getString(R.string.prifi_config_shared_preferences), MODE_PRIVATE).edit();
            editor.putString(getString(R.string.prifi_client_public_key), pubKey);
            editor.putString(getString(R.string.prifi_client_private_key), priKey);
            editor.putBoolean(getString(R.string.prifi_config_first_init), false);
            editor.apply();
        } else {
            final String currentPubKey = preferences.getString(getString(R.string.prifi_client_public_key),"");
            final String currentPriKey = preferences.getString(getString(R.string.prifi_client_private_key),"");

            try {
                Servicego.setPublicKey(currentPubKey);
                Servicego.setPrivateKey(currentPriKey);

            } catch (Exception e) {
                throw new RuntimeException("Can't set service parameters");
            }
        }
    }

}
