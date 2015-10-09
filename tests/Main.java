import java.util.Scanner;

public class Main{

    public static void main(String[] args){
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int sub = 0;
        StringBuilder stringBuilder;
        stringBuilder = new StringBuilder() ;
        for( int i=0 ; i < n ; i++){

            stringBuilder.append(  scanner.next()) ;
            stringBuilder.delete(0,stringBuilder.lastIndexOf("@")+1);
            while( stringBuilder.length()>0 &&  stringBuilder.charAt(0) == '#'){
               stringBuilder.deleteCharAt(0);
            }
            while( stringBuilder.length()>0 &&(sub = stringBuilder.indexOf("#")) >0){
                stringBuilder =  stringBuilder.delete(sub-1,sub+1);
            }
            if( stringBuilder.length() == 0){
                System.out.println("");
            }else
            System.out.println( stringBuilder.toString());
            stringBuilder.delete(0,stringBuilder.length());
        }
    }

}