#include "stm32f10x.h"                  // Device header
#include "Tipkalo.h"                  // Device header



void GPIO_ClockCtrl(GPIO_TypeDef *pGPIOx, uint8_t EnOrDi){
if(EnOrDi==ENABLE){
	if(pGPIOx==GPIOA){RCC->APB2ENR |= RCC_APB2ENR_IOPAEN ;
}else if(pGPIOx==GPIOB){
RCC->APB2ENR |= RCC_APB2ENR_IOPBEN ;
}else if(pGPIOx==GPIOC){
RCC->APB2ENR |= RCC_APB2ENR_IOPCEN ;
}}
if(EnOrDi==DISABLE){
	if(pGPIOx==GPIOA){RCC->APB2ENR &=~ RCC_APB2ENR_IOPAEN ;
}else if(pGPIOx==GPIOB){
RCC->APB2ENR &=~ RCC_APB2ENR_IOPBEN ;
}else if(pGPIOx==GPIOC){
RCC->APB2ENR &=~ RCC_APB2ENR_IOPCEN ;
}}
}

void GPIO_PinMode(GPIO_TypeDef *pGPIOx, uint16_t pinNumber, uint16_t pinMode, uint32_t pinSpeed=FAST){
	if(pinMode==INPUT){
	pGPIOx->CRL &=~ (1<<pinNumber*4);	
	pGPIOx->CRL &=~ (1<<(pinNumber*4)+1);
	pGPIOx->CRL &=~ (1<<(pinNumber*4)+2);	
	pGPIOx->CRL  |= (1<<(pinNumber*4)+3);
	}else
	if(pinMode==OUTPUT){
	pGPIOx->CRL |= (0<<pinNumber*4);	
	pGPIOx->CRL |= (1<<pinNumber*4+1);
	pGPIOx->CRL &=~(1<<pinNumber*4+2);	
	pGPIOx->CRL &=~(1<<pinNumber*4+3);
	}
	


}
void GPIO_DigitalWrite(GPIO_TypeDef *pGPIOx,uint16_t PinNumber, bool Value){
		
	if(Value==HIGH){
	pGPIOx->ODR |=(Value<<PinNumber);
	
	}
if(Value==LOW){
	pGPIOx->ODR &=~(1<<PinNumber);
	
	}

	

}void daly(int n){
		for(int i; i<n ; i++){;}
		
		}
		bool GPIO_DigitalRead(GPIO_TypeDef *pGPIOx,uint16_t pinNumber){
		
		bool ulaz=false;
			if(pinNumber==0)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR0;
			if(pinNumber==1)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR1;
			if(pinNumber==2)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR2;
			if(pinNumber==3)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR3;
			if(pinNumber==4)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR4;
			if(pinNumber==5)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR5;
			if(pinNumber==6)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR6;
			if(pinNumber==7)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR7;
			if(pinNumber==8)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR8;
 			if(pinNumber==9)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR9;
			if(pinNumber==10)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR10;
			if(pinNumber==11)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR11;
			if(pinNumber==12)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR12;
			if(pinNumber==13)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR13;
			if(pinNumber==14)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR14;
			if(pinNumber==15)
			ulaz = pGPIOx->IDR & GPIO_IDR_IDR15;
			
		return ulaz;
		}
			void GPIO_DigitalWritePort(GPIO_TypeDef *pGPIOx,uint16_t Value){
			
				pGPIOx->ODR |= Value;
			
			
			}
		
uint16_t GPIO_DigitalReadPort(GPIO_TypeDef *pGPIOx){

uint16_t ulaz;
ulaz= pGPIOx->ODR;

return ulaz;
}		

void GPIO_ToggleOutputPin(GPIO_TypeDef *pGPIOx,uint16_t pinNumber){

			if(pinNumber==0)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR0)<<pinNumber);
			if(pinNumber==1)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR1)<<pinNumber);
			if(pinNumber==2)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR2)<<pinNumber);
			if(pinNumber==3)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR3)<<pinNumber);
			if(pinNumber==4)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR4)<<pinNumber);
			if(pinNumber==5)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR5)<<pinNumber);
			if(pinNumber==6)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR6)<<pinNumber);
			if(pinNumber==7)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR7)<<pinNumber);
			if(pinNumber==8)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR8)<<pinNumber);
			if(pinNumber==9)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR9)<<pinNumber);
			if(pinNumber==10)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR10)<<pinNumber);
			if(pinNumber==11)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR11)<<pinNumber);
			if(pinNumber==12)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR12)<<pinNumber);
			if(pinNumber==13)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR13)<<pinNumber);
			if(pinNumber==14)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR14)<<pinNumber);
			if(pinNumber==15)
				pGPIOx->ODR |=(!(pGPIOx->ODR && GPIO_ODR_ODR15)<<pinNumber);
}
