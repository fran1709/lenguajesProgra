package com;

public class Operador {
    // Atributos
    private String signo = "";

    // Constructor
    public Operador(String pOperador){
        this.signo = pOperador;
    }

    //Methos
    public String getSigno() {
        return signo;
    }
    public void setSigno(String signo) {
        this.signo = signo;
    }

    @Override
    public String toString() {
        String text = "";
        text += this.getSigno();
        return text;
    }
}
