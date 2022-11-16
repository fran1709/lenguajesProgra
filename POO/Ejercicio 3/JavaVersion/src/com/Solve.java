package com;

import java.util.ArrayList;
import java.util.StringTokenizer;

public class Solve {
    // Constructor
    public Solve(){}

    // Methods
    /**
     * Método encargado de splitear los objetos.
     * @param pOperation
     * @return
     */
    private ArrayList<String> enlistar(String pOperation) {
        ArrayList<String> datos = new ArrayList<String>();
        StringTokenizer tokens = new StringTokenizer(pOperation, "[*/+-]", true);
        while(tokens.hasMoreTokens()){
            datos.add(tokens.nextToken());
        }
        return datos;
    }
    /**
     * Método encargado de convertir un String->Double
     * @param pNum
     * @return
     */
    private double convert_numero(String pNum){
        double num = 0.0;
        try{num = Double.parseDouble(pNum);}
        catch(NumberFormatException e){System.out.println("Formato inválido -> " + e);}
        return num;
    }
    /**
     * Método encargado de verificar si un String es una operador.
     * @param pValue
     * @return
     */
    private boolean isOp(String pValue){
        boolean isOp = true;
        if (pValue.equals("*") || pValue.equals("/") || 
            pValue.equals("+") || pValue.equals("-")) {
            isOp = false;
        }
        return isOp;
    }
    /**
     * Método encargado de convertir los strings en Objetos Operador || Numero
     * @param pValues
     * @param pObjetos
     */
    private void instanciar(ArrayList<String> pValues,  ArrayList<Object> pObjetos) {
        for (int i=0; i<pValues.size(); i++){
            if(!isOp(pValues.get(i))){
                pObjetos.add(new Operador(pValues.get(i)));
            } else {
                pObjetos.add(new Numero(convert_numero(pValues.get(i))));
            }
        }
    }
    /**
     * Método encargado de calcular el valor resultante mediante operaciones matemáticas.
     * @param pValues
     * @return result double 
     */
    private double calcular(ArrayList<Object> pValues) {
        double result = 0.0;
        for (int i=0; i<pValues.size(); i++){
            if(pValues.get(i).toString().equals("*")){
                result *= Double.parseDouble(pValues.get(i+1).toString());
                //System.out.println("0 = "+ pValues.get(i+1).toString());
                i+=1;
            } else if(pValues.get(i).toString().equals("/")){
                result /= Double.parseDouble(pValues.get(i+1).toString());
                //System.out.println("1 = " + pValues.get(i+1).toString());
                i+=1;
            } else if(pValues.get(i).toString().equals("+")){
                result += Double.parseDouble(pValues.get(i+1).toString());
                //System.out.println("2 = " + pValues.get(i+1).toString());
                i+=1;
            } else if ( (pValues.get(i).toString()).equals("-")){
                result -= Double.parseDouble(pValues.get(i+1).toString());
                //System.out.println("3 = " + pValues.get(i+1).toString());
                i+=1;
            } else {
                //System.out.println("4 = " + pValues.get(i+1).toString());
                result = Double.parseDouble(pValues.get(i).toString());
            }
            //System.out.println("Valor -> " + result);
        }
        return result;
    }
    /**
     * Método encargado del procedimiento.
     * @param pOperation
     * @return
     */
    public double result(String pOperation) {
        ArrayList<Object> objetos = new ArrayList<Object>();
        ArrayList<String> datos = enlistar(pOperation);
        instanciar(datos, objetos);
        return calcular(objetos);
    }
}
