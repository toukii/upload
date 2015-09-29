package t360;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Main{
    public void Flower(int a,int b){
        List<Integer> l = new ArrayList<Integer>();
        int n,sum;
    	for(int i = a;i<=b;i++){
        	sum = 0;
            n = i;
            while(n>0){
            	sum += Math.pow(n%10,3);
                n /=10;
            }
            if(sum == i) l.add(i);
        }
        int size = l.size();
        if (size<=0) {
			System.out.println("no");
			return;
		}
        for(int i=0;i<size-1;i++){
        	System.out.print(l.get(i)+" ");
        }
        System.out.print(l.get(size-1));
    }
	public static void main(String[] as){
		Main m = new Main();
    	Scanner sc = new Scanner(System.in);
        int a,b;
        while(sc.hasNext()){
        	a = sc.nextInt();
            b = sc.nextInt();
            m.Flower(a,b);
        }
    }
}