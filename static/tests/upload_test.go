import java.util.Scanner;

public class Main {
	public void Fire(Scanner sc,long n, long a) {
		float sum = a;
		long p = 0;
		for (long i = 0; i < n; i++) {
			p = sc.nextLong();
			System.out.println(p);
			if (p<=sum) {
				sum += p;
			}else{
				sum += Main.gcd(sum,p);
			}
		}
		System.out.printf("%.0f\n", sum);
	}
	public static double gcd(double a, double b) {
		if(a<b){
            double temp;
            temp=a;
            a=b;
            b=temp;	
            /*a=a^b;
			 b=a^b;
			 a=a^b;*/
	    }
	    if(0==b){
	        return a;
	    }
	    return gcd(a-b,b);
	}
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while(sc.hasNext()){
	        long n = sc.nextLong();
	        long a = sc.nextLong();
	        Main m = new Main();
	        m.Fire(sc, n, a);
        }
    }
}
