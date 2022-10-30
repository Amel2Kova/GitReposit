using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zadatak_5
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void cmdIzracunaj_Click(object sender, EventArgs e)
        {
            double prvi, drugi;
            double rijesenje;
            //prvi = txtPrvi.Text;

            prvi = double.Parse(txtPrvi.Text);
            drugi = double.Parse(txtDrugi.Text);


            rijesenje = prvi + drugi;
            MessageBox.Show(rijesenje.ToString("F2"));


        }

        private void txtPrvi_TextChanged(object sender, EventArgs e)
        {

        }
    }
}
