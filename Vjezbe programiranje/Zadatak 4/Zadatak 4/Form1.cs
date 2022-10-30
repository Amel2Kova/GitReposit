using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace Zadatak_4
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void cmdLabel_Click(object sender, EventArgs e)
        {
            cmdLabel.BackColor = Color.Yellow;
            cmdLabel.ForeColor = Color.Blue;
            cmdLabel.Text = "Osnovne promjene";
            cmdLabel.Location = new Point(222, 90);
            cmdLabel.Size = new Size(300, 40);
            cmdLabel.BorderStyle = BorderStyle.FixedSingle;
            cmdLabel.Font = new Font("Arial", 40);
            

            
                
        }
    }
}
