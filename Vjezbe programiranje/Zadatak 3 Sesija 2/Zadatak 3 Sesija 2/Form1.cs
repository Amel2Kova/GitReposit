using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zadatak_3_Sesija_2
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            string Odabir= "Odabrali ste :  ";
            if (radioButton1.Checked)
                Odabir = Odabir + radioButton1.Text;

            if (radioButton2.Checked)
                Odabir = Odabir + radioButton2.Text;

            if (radioButton3.Checked)
                Odabir = Odabir + radioButton3.Text;

            textBox1.Text = Odabir;
        }

        
       
    }
}
