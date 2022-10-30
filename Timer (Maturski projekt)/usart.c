#include "usart.h"

//Функция инициализации периферии
void Init_usart()
{
  //Включаем тактирование GPIOA, USART1 и альтернативных функций AFIO
  RCC_APB2PeriphClockCmd((RCC_APB2Periph_USART1 | RCC_APB2Periph_GPIOA | RCC_APB2Periph_AFIO), ENABLE);

  //Инициализации вывода PA9 - USART1_Tx
  GPIO_InitStruct.GPIO_Pin = GPIO_Pin_9; //Настройки вывода PA9
  GPIO_InitStruct.GPIO_Speed = GPIO_Speed_50MHz; //Скорость порта максимальная
  GPIO_InitStruct.GPIO_Mode = GPIO_Mode_AF_PP; //Режим альтернативной функции, выход Push-Pull
  GPIO_Init(GPIOA, &GPIO_InitStruct); //Заданные настройки сохраняем в регистрах GPIOА

  //Инициализации вывода PA10 - USART1_Rx
  GPIO_InitStruct.GPIO_Pin = GPIO_Pin_10; //Настройки вывода PA10
  GPIO_InitStruct.GPIO_Mode  = GPIO_Mode_IN_FLOATING; //Input floating
  GPIO_Init(GPIOA, &GPIO_InitStruct); //Заданные настройки сохраняем в регистрах GPIOА


  //Инициализация USART1
  USART_InitStruct.USART_BaudRate = 115200; //Скорость обмена 9600 бод
  USART_InitStruct.USART_WordLength = USART_WordLength_8b; //Длина слова 8 бит
  USART_InitStruct.USART_StopBits = USART_StopBits_1; //1 стоп-бит
  USART_InitStruct.USART_Parity = USART_Parity_No ; //Без проверки четности
  USART_InitStruct.USART_HardwareFlowControl = USART_HardwareFlowControl_None; //Без аппаратного контроля
  USART_InitStruct.USART_Mode = USART_Mode_Rx | USART_Mode_Tx; //Включен передатчик и приемник USART1
  USART_Init(USART1, &USART_InitStruct); //Заданные настройки сохраняем в регистрах USART1

  USART_Cmd(USART1, ENABLE); //Включаем USART1
}

//Функция передачи символа
void Usart1_Send_symbol(uint8_t data)
{
  while(!(USART1->SR & USART_SR_TC)); //Проверяем установку флага TC - завершения предыдущей передачи
  USART1->DR = data; //Записываем значение в регистр данных - передаем символ
}
