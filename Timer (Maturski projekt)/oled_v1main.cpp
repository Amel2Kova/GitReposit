#include "stm32f10x.h"                  // Device header
#include <stdio.h>
#include "stdlib.h"
#include "ssd1306.h"
#include "Tipkalo.h"  


void testString(int x);
int radius,posx,posy,j,x;
char* int2char(int iNumber);
int main(void)
{
GPIO_ClockCtrl(GPIOA,ENABLE);

SSD1306_Init();


	
	while(1)
	{
	int n=59;
				int m=59;
		int x;
	j>=0? j--:j=n;
			if(j==0){x>0?x--:x=m;}
		
		//----------------hours
		GPIO_PinMode(GPIOA,0,INPUT,FAST);
		
		GPIO_PinMode(GPIOA,1,INPUT,FAST);
		
		GPIO_PinMode(GPIOA,2,INPUT,FAST);
		
		GPIO_PinMode(GPIOA,3,INPUT,FAST);
		
		bool a0=false,a1=false,a2=false,a3=false;
		
		a0=GPIO_DigitalRead(GPIOA,0);
		a1=GPIO_DigitalRead(GPIOA,1);
		a2=GPIO_DigitalRead(GPIOA,2);
		a3=GPIO_DigitalRead(GPIOA,3);
		SSD1306_GotoXY( 15,0);
			SSD1306_Puts("Smajaa!! Timer", &Font_7x10, SSD1306_COLOR_WHITE);
		SSD1306_GotoXY( 0,18);
			SSD1306_Puts("------------------", &Font_7x10, SSD1306_COLOR_WHITE);
	
			if(a0==true)
		{
		while(1)
		{
			bool Pritisnuto,P2;
		a1=GPIO_DigitalRead(GPIOA,1);
		a2=GPIO_DigitalRead(GPIOA,2);
		a3=GPIO_DigitalRead(GPIOA,3);
		if(a1==true){
		 
			if(!Pritisnuto)
			x++;else if(59<x)x=0;
			Pritisnuto=true;
			
		}else  {Pritisnuto=false;
		}
		if(a2==true)
		{
			if(x>0&&!P2)
			x--;else if(x==0)x=59; 
			P2=true;
			
		
		}else {P2=false;
		}
			SSD1306_GotoXY(18 ,35);
		testString(x);
		SSD1306_UpdateScreen(); 
		
		
		if(a3==true)
		{
		break;
		}
		}
		
		}
		SSD1306_GotoXY( 18,35);
		testString(x);
				SSD1306_GotoXY( 54,35);
		SSD1306_Puts(":", &Font_16x26, SSD1306_COLOR_WHITE);
			SSD1306_GotoXY(64,35);
					testString(j);

		    
					//testString(0);
					//SSD1306_Puts("00", &Font_16x26, SSD1306_COLOR_WHITE);
			    //SSD1306_DrawRectangle(132-i,36,8,8,SSD1306_COLOR_WHITE);

					SSD1306_UpdateScreen(); 
			           for(int i=0;i<10000;i++)
		{}
	

			
	}

}
char* int2char(int iNumber){
    int iNumbersCount=0;
    int iTmpNum=iNumber;
    while(iTmpNum){
        iTmpNum/=10;
        iNumbersCount++;
    }
    char *buffer=new char[iNumbersCount+1];
    for(int i=iNumbersCount-1;i>=0;i--){
        buffer[i]=(char)((iNumber%10)|48);
        iNumber/=10;
    }
    buffer[iNumbersCount]='\0';
    return buffer;

}

void testString(int x)
{
	switch(x){
		case 0:
	   SSD1306_Puts("00", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
		case 1:
	   SSD1306_Puts("01", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 2:
	   SSD1306_Puts("02", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
						case 3:
	   SSD1306_Puts("03", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
								case 4:
	   SSD1306_Puts("04", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 5:
	   SSD1306_Puts("05", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 6:
	   SSD1306_Puts("06", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 7:
	   SSD1306_Puts("07", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 8:
	   SSD1306_Puts("08", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
						case 9:
	   SSD1306_Puts("09", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
								case 10:
	   SSD1306_Puts("10", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 11:
	   SSD1306_Puts("11", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 12:
	   SSD1306_Puts("12", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 13:
	   SSD1306_Puts("13", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 14:
	   SSD1306_Puts("14", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
						case 15:
	   SSD1306_Puts("15", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
								case 16:
	   SSD1306_Puts("16", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 17:
	   SSD1306_Puts("17", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 18:
	   SSD1306_Puts("18", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 19:
	   SSD1306_Puts("19", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
						case 20:
	   SSD1306_Puts("20", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 21:
	   SSD1306_Puts("21", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
						case 22:
	   SSD1306_Puts("22", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
								case 23:
	   SSD1306_Puts("23", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 24:
	   SSD1306_Puts("24", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 25:
	   SSD1306_Puts("25", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
	  		case 26:
	   SSD1306_Puts("26", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 27:
	   SSD1306_Puts("27", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
						case 28:
	   SSD1306_Puts("28", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
								case 29:
	   SSD1306_Puts("29", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 30:
	   SSD1306_Puts("30", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 31:
	   SSD1306_Puts("31", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 32:
	   SSD1306_Puts("32", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 33:
	   SSD1306_Puts("33", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
						case 34:
	   SSD1306_Puts("34", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
								case 35:
	   SSD1306_Puts("35", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 36:
	   SSD1306_Puts("36", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 37:
	   SSD1306_Puts("37", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 38:
	   SSD1306_Puts("38", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 39:
	   SSD1306_Puts("39", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 40:
	   SSD1306_Puts("40", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
						case 41:
	   SSD1306_Puts("41", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
								case 42:
	   SSD1306_Puts("42", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 43:
	   SSD1306_Puts("43", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 44:
	   SSD1306_Puts("44", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
																						case 45:
	   SSD1306_Puts("45", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 46:
	   SSD1306_Puts("46", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 47:
	   SSD1306_Puts("47", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
				case 48:
	   SSD1306_Puts("48", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
						case 49:
	   SSD1306_Puts("49", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
								case 50:
	   SSD1306_Puts("50", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 51:
	   SSD1306_Puts("51", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 52:
	   SSD1306_Puts("52", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
																				case 53:
	   SSD1306_Puts("53", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 54:
	   SSD1306_Puts("54", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 55:
	   SSD1306_Puts("55", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
																				case 56:
	   SSD1306_Puts("56", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
										case 57:
	   SSD1306_Puts("57", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												case 58:
	   SSD1306_Puts("58", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
																			case 59:
	   SSD1306_Puts("59", &Font_16x26, SSD1306_COLOR_WHITE);
    break;
												
		default:
			SSD1306_Puts("00", &Font_16x26, SSD1306_COLOR_WHITE);
	}

}
