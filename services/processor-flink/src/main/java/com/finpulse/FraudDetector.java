package com.finpulse;

import org.apache.flink.api.common.functions.FilterFunction;
import org.apache.flink.streaming.api.datastream.DataStream;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

public class FraudDetector {
    public static void main(String[] args) throws Exception {
        final StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();
        ObjectMapper tempMapper = new ObjectMapper();

        // Simulate consuming from Kinesis/Kafka (as used in your Amazon & Vanguard roles)
        DataStream<String> rawStream = env.fromElements(
            "{\"transaction_id\": \"1\", \"amount\": 95000.0, \"type\": \"TRANSFER\"}",
            "{\"transaction_id\": \"2\", \"amount\": 10.0, \"type\": \"DEBIT\"}"
        );

        DataStream<String> alerts = rawStream.filter(new FilterFunction<String>() {
            @Override
            public boolean filter(String value) throws Exception {
                JsonNode node = tempMapper.readTree(value);
                // Anomaly detection logic: Flag transactions over $50,000
                return node.get("amount").asDouble() > 50000.0;
            }
        });

        alerts.print(); // In production, this would sink to DynamoDB or SNS [cite: 28, 61]
        env.execute("FinPulse Real-Time Fraud Detection");
    }
}