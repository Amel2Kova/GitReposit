using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zadatak_5_Sesija_2
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        double Var1, Var2, Rez;
        string odabir;
        private void button1_Click(object sender, EventArgs e)
        {
            Var1 = double.Parse(textBox1.Text);
            odabir = button1.Text;
            textBox1.Text = "";
        }
        private void button2_Click(object sender, EventArgs e)
        {
            Var1 = double.Parse(textBox1.Text);
            odabir = button2.Text;
            textBox1.Text = "";
        }
        private void button3_Click(object sender, EventArgs e)
        {
            Var1 = double.Parse(textBox1.Text);
            odabir = button3.Text;
            textBox1.Text = "";
        }

        private void button4_Click(object sender, EventArgs e)
        {
            Var2 = double.Parse(textBox1.Text);
            if (odabir == "+")
                Rez = Var1 + Var2;
            if (odabir == "-")
                Rez = Var1 - Var2;
            if (odabir == "*")
                Rez = Var1 * Var2;
            if (odabir == "/")
                Rez = Var1 / Var2;
            textBox1.Text = Convert.ToString(Rez);
            
        }

        private void button5_Click(object sender, EventArgs e)
        {
            Var1 = double.Parse(textBox1.Text);
            odabir = button5.Text;
            textBox1.Text = "";
        }

        

        
       
    }
}
