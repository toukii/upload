import java.util.Scanner;

public class Main {
	public void fire(int iR,int[][] input) {
		float R = iR*iR;
		float r = 0;
		int count = 0;
		int a = input[3][0];
		int b = input[3][1];
		for (int i = 0; i < 3; i++) {
			r = 0;
			r = (float)(Math.pow(input[i][0]-a,2)+Math.pow(input[i][1]-b,2));
			if (r<=R) {
				count++;
			}
		}
		System.out.printf("%dX",count);
	}
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int R = sc.nextInt();
        int[][] input = new int[4][2];
        for (int i = 0; i < 4; i++) {
			input[i][0] = sc.nextInt();
			input[i][1] = sc.nextInt();
		}
        Main m = new Main();
        m.fire(R, input);
    }
}
