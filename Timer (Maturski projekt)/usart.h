#include "stm32f10x.h"

//Структуры для инициализации GPIOA и USART1
GPIO_InitTypeDef    GPIO_InitStruct;
USART_InitTypeDef    USART_InitStruct;

void Init_usart(void); //Объявление функции инициализации периферии
void Usart1_Send_symbol(uint8_t); //Объявление функции передачи символа
void Usart1_Send_String(char* str); //Объявление функции передачи строки

char* str;
uint8_t data;
