//
//  360.h
//  offer
//
//  Created by 郭浩 on 15/9/29.
//  Copyright (c) 2015年 Guohao. All rights reserved.
//

#ifndef offer__60_h
#define offer__60_h

#include<string>
#include<deque>
#include<iostream>
#include<vector>

using namespace std;

string convert(const string& str){
    deque<char> line;
    if(str.length()>0){
        int len = (int)str.length();
        for(int i=0;i<len;i++){
            if(str[i]=='#'){
                if(!line.empty())
                    line.pop_back();
                //                else
                //                    return "";
            }
            else if (str[i] == '@'){
                if(!line.empty()){
                    while(!line.empty())
                        line.pop_back();
                }
                //                else
                //                    return "";
            }
            else{
                line.push_back(str[i]);
            }
        }
    }
    if(!line.empty()){
        string str(line.begin(), line.end());
        return str;
    }
    return "";
}

int _main(int argc, const char * argv[]) {
    vector<string> code;
    int m;
    cin>>m;
    string str;
    while(m>0){
        cin>>str;
        code.push_back(convert(str));
        m--;
    }
    vector<string>::iterator it;
    for(it = code.begin();it!=code.end();it++)
        cout<<*it<<endl;
    return 0;
}

int _main2(){
    int m,n,a,b;
    vector<int>ret;
    while(true){
        cin>>m>>n;
        if(m!=0&&n!=0){
            while(n>0){
                cin>>a>>b;
                
            }
        }
    }
}

#endif
