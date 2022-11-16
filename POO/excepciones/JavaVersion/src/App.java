import com.Solve;

public class App {
    public static void main(String[] args) {
        Solve Calculadora = new Solve();
        String operacion = "512*30/2+320-1";
        System.out.println("Resultado -> " + Calculadora.result(operacion)); 
    }
}
