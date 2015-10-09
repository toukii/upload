import java.util.Scanner;

public class Main {
	public void count(long n) {
		long k = (int)Math.sqrt(2*n);
		if (k*(k+1)>2*n) {
			k -=1;
		}
		double sum = 0;
		sum = 1 + (k-2)*(k-2+1)/2;
		sum += n - (k*(k+1)/2);
		System.out.printf("%.0f\n", sum);
		
	}
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while(sc.hasNext()){
	        long n = sc.nextInt();
	        Main m = new Main();
	        m.count(n);
        }
    }
}
