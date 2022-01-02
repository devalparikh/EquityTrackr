package com.EquityTrackr.backend.investor;

public class Investor {
    private String name;
    private String email;
    private Long balance;

    public Investor() {
    }

    public Investor(String name, String email) {
        this.name = name;
        this.email = email;
        this.balance = (long) 0;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public Long getBalance() {
        return balance;
    }

    public void setBalance(Long balance) {
        this.balance = balance;
    }

    @Override
    public String toString() {
        return "Investor{" +
                ", name='" + name + '\'' +
                ", email='" + email + '\'' +
                '}';
    }
}
