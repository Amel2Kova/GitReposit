using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace zadatak1
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }



        private void checkBox4_CheckedChanged(object sender, EventArgs e)
        {
            if (textBox1.Visible == true)
                textBox1.Visible = false;
            else
                textBox1.Visible = true;
        }

        private void button1_Click(object sender, EventArgs e)
        {

            string Odabir = "Odabrali ste:" + Environment.NewLine;
            if (Dorucak.Checked)
                Odabir = Odabir + Dorucak.Text+Environment.NewLine;

            if (Rucak.Checked)
                Odabir = Odabir + Rucak.Text + Environment.NewLine;

            if (Vecera.Checked)
                Odabir = Odabir + Vecera.Text + Environment.NewLine;
            textBox1.Text = (Odabir);
        }

       
        }

      

    }

       
       

       

       
    

