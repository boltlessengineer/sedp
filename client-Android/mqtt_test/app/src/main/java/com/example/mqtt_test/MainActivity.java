package com.example.mqtt_test;

import androidx.appcompat.app.AppCompatActivity;
import androidx.core.app.NotificationCompat;
import androidx.core.app.NotificationManagerCompat;

import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.content.Intent;
import android.media.Ringtone;
import android.media.RingtoneManager;
import android.net.Uri;
import android.os.Build;
import android.os.Bundle;
import android.os.Vibrator;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.CompoundButton;
import android.widget.EditText;
import android.widget.Switch;
import android.widget.TextView;
import android.widget.Toast;

import org.eclipse.paho.android.service.MqttAndroidClient;
import org.eclipse.paho.client.mqttv3.IMqttActionListener;
import org.eclipse.paho.client.mqttv3.IMqttDeliveryToken;
import org.eclipse.paho.client.mqttv3.IMqttToken;
import org.eclipse.paho.client.mqttv3.MqttCallback;
import org.eclipse.paho.client.mqttv3.MqttClient;
import org.eclipse.paho.client.mqttv3.MqttException;
import org.eclipse.paho.client.mqttv3.MqttMessage;

public class MainActivity extends AppCompatActivity {

    public String MQTTHOST = "";
    EditText IP_address;

    //public String USERNAME = "username";
    //public String PASSWORD = "password";

    public String pubTopic = "test";
    EditText subTopic;

    MqttAndroidClient client;

    TextView subText;
    EditText pubText;

    Button subButton;

    //MqttConnectOptions options;

    Vibrator vibrator;
    Ringtone myRingtone;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        pubText = (EditText)findViewById(R.id.pubText);
        IP_address = (EditText)findViewById(R.id.IP_address);
        subTopic = (EditText)findViewById(R.id.subTopic);

        subText = (TextView)findViewById(R.id.subText); //subText add

        vibrator = (Vibrator) getSystemService(VIBRATOR_SERVICE);
        Uri uri = RingtoneManager.getDefaultUri(RingtoneManager.TYPE_NOTIFICATION);
        myRingtone = RingtoneManager.getRingtone(getApplicationContext(), uri);

        subButton = (Button)findViewById(R.id.subButton);

        subButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                setSubscription();
            }
        });

        setHost();

        try {
            IMqttToken token = client.connect();
            //IMqttToken token = client.connect(options);
            token.setActionCallback(new IMqttActionListener() {
                @Override
                public void onSuccess(IMqttToken asyncActionToken) {
                    // We are connected
                    //Log.d(TAG, "onSuccess");
                    Toast.makeText(MainActivity.this, "connected", Toast.LENGTH_SHORT).show();
                    setSubscription();
                }

                @Override
                public void onFailure(IMqttToken asyncActionToken, Throwable exception) {
                    // Something went wrong e.g. connection timeout or firewall problems
                    //Log.d(TAG, "onFailure");
                    Toast.makeText(MainActivity.this, "failed to connect", Toast.LENGTH_SHORT).show();

                }
            });
        } catch (MqttException e) {
            e.printStackTrace();
        }
        catch (Exception e) {
            e.printStackTrace();
        }

        client.setCallback(new MqttCallback() { //setCallback for message arrive
            @Override
            public void connectionLost(Throwable cause) {

            }

            @Override
            public void messageArrived(String topic, MqttMessage message) throws Exception {
                subText.setText(new String(message.getPayload())); //show text on subText

                String text = new String(message.toString());
                Log.d("valueOf", String.valueOf(message.getPayload()));
                Log.d("text", text);

                if(text.equals("hello"))
                {
                    Intent intent = new Intent(MainActivity.this, MainActivity.class);
                    PendingIntent pendingIntent = PendingIntent.getActivity(MainActivity.this, 0, intent,PendingIntent.FLAG_UPDATE_CURRENT);

                    if(Build.VERSION.SDK_INT >= Build.VERSION_CODES.O)
                    {
                        NotificationChannel channel = new NotificationChannel("notify test", "notiName", NotificationManager.IMPORTANCE_DEFAULT);
                        NotificationManager manager = getSystemService(NotificationManager.class);
                        manager.createNotificationChannel(channel);
                    }

                    NotificationCompat.Builder builder = new NotificationCompat.Builder(MainActivity.this, "notify test");
                    builder.setContentTitle("알림 테스트")
                            .setContentText("이것은 알림 테스트입니다.")
                            .setSmallIcon(R.drawable.ic_alarm_icon)
                            .setDefaults(Notification.DEFAULT_VIBRATE)
                            .setContentIntent(pendingIntent)
                            .setAutoCancel(true);

                    NotificationManagerCompat managerCompat = NotificationManagerCompat.from(MainActivity.this);
                    managerCompat.notify(1,builder.build());

                    Toast.makeText(MainActivity.this, "알림 생성", Toast.LENGTH_SHORT).show();
                }
            }

            @Override
            public void deliveryComplete(IMqttDeliveryToken token) {

            }
        });
    }

    public void pub(View v)
    {
        String topic = pubTopic;
        String message = pubText.getText().toString();
        Log.d("message", message);
        try {
            client.publish(topic, message.getBytes(), 0, false);
        } catch (MqttException e) {
            e.printStackTrace();
        }
    }

    private void setHost()
    {
        String _HOST = IP_address.getText().toString();
        MQTTHOST = "tcp://"+_HOST+":1883";
        Log.d("HOST", MQTTHOST);
        String clientId = MqttClient.generateClientId();
        client = new MqttAndroidClient(this.getApplicationContext(), MQTTHOST, clientId);

        //options = new MqttConnectOptions();
        //options.setUserName(USERNAME);
        //options.setPassword(PASSWORD.toCharArray());
    }

    public void setSubscription() //subscribe for TOPIC
    {
        try
        {
            String topic = subTopic.getText().toString();
            Log.d("subTopic", topic);
            client.subscribe(topic, 0);
            Toast.makeText(MainActivity.this, "subscribe "+ topic, Toast.LENGTH_SHORT).show();
        }
        catch(MqttException e)
        {
            e.printStackTrace();
        }
    }

    public void conn(View v) //try connecting to MQTT
    {
        try {
            setHost();

            //Intent intent = new Intent(MainActivity.this, notificationService.class);
            //startService(intent);

            //Toast.makeText(MainActivity.this, "Service start", Toast.LENGTH_SHORT).show();

            IMqttToken token = client.connect();
            //IMqttToken token = client.connect(options);
            token.setActionCallback(new IMqttActionListener() {
                @Override
                public void onSuccess(IMqttToken asyncActionToken) {
                    // We are connected
                    //Log.d(TAG, "onSuccess");
                    Toast.makeText(MainActivity.this, "connected", Toast.LENGTH_SHORT).show();
                    setSubscription();
                }

                @Override
                public void onFailure(IMqttToken asyncActionToken, Throwable exception) {
                    // Something went wrong e.g. connection timeout or firewall problems
                    //Log.d(TAG, "onFailure");
                    Toast.makeText(MainActivity.this, "failed to connect", Toast.LENGTH_SHORT).show();

                }
            });
        } catch (MqttException e) {
            e.printStackTrace();
        }
    }

    public void disconn(View v) //disconnect MQTT
    {
        try {

            //Intent intent = new Intent(MainActivity.this, notificationService.class);
            //stopService(intent);

            //Toast.makeText(MainActivity.this, "Service stop", Toast.LENGTH_SHORT).show();

            IMqttToken token = client.disconnect();
            token.setActionCallback(new IMqttActionListener() {
                @Override
                public void onSuccess(IMqttToken asyncActionToken) {
                    // We are connected
                    //Log.d(TAG, "onSuccess");
                    Toast.makeText(MainActivity.this, "disconnected", Toast.LENGTH_SHORT).show();
                }

                @Override
                public void onFailure(IMqttToken asyncActionToken, Throwable exception) {
                    // Something went wrong e.g. connection timeout or firewall problems
                    //Log.d(TAG, "onFailure");
                    Toast.makeText(MainActivity.this, "could not disconnect..", Toast.LENGTH_SHORT).show();

                }
            });
        } catch (MqttException e) {
            e.printStackTrace();
        }
    }

}