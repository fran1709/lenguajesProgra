package com;

public class Numero {
    // Atributos
    private double valor = 0;

    // Constructor
    public Numero(double pNum){
        this.valor = pNum;
    }

    // Methods
    public double getValor() {
        return valor;
    }
    public void setValor(double valor) {
        this.valor = valor;
    }
    
    @Override
    public String toString() {
        String text = "";
        text += this.getValor();
        return text;
    }
}
