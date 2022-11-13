package elements.expresion;

public abstract class Expresion {
    private String expres;

    public Expresion(String pExpresion){
        this.expres = pExpresion;
    }

    public String getExpres() {
        return expres;
    }

    public void setExpres(String expres) {
        this.expres = expres;
    }
    
}
