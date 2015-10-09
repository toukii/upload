

import java.util.Scanner;
/*Third*/
public class Second {
	boolean[] visit;
	int [][]input;
	int n;
	int N;
	public Second(int[][] input, int N,int n){
		this.N = N;
		this.n = n;
		this.input = input;
		visit = new boolean[N];
		for (int i = 0; i < N; i++) {
			visit[i] = false;
		}
	}
	public void Count() {
		visitNode(1);
		int count = 0;
		for (boolean bi : visit) {
			if (bi) {
				count ++;
			}
		}
		count --;
		if (count>0) {
			System.out.println(count);
		}else{
			System.out.println(0);
		}
	}
	public void visitNode(int target){
		if (target<=0) {
			return;
		}
		if (visit[target-1]) {
			return;
		}
		visit[target-1] = true;
		for (int i = 0; i < n; i++) {
			if (input[i][0]==target) {
				visitNode(input[i][1]);
			}else if (input[i][1]==target) {
				visitNode(input[i][0]);
			}
		}
		
	}
	public static void main(String[] args) {
    	Scanner sc = new Scanner(System.in);
        int N,n;
        Second m;
        int input [][] ;
        while(sc.hasNext()){
        	N = sc.nextInt();
            n = sc.nextInt();
            if (N==0) {
				return;
			}
            input = new int[n][2];
            /*for (int i = 0; i < N; i++) {
				input[i]= new int[2]; 
			}*/
            for (int i = 0; i < n; i++) {
				input[i][0] = sc.nextInt();
				input[i][1] = sc.nextInt();
			}

    		m = new Second(input,N,n);
            m.Count();
        }
	}

}
