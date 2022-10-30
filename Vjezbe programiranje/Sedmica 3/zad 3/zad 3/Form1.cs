using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace zad_3
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();

              
            checkedListBox1.Items.Add("Nedelja", CheckState.Checked);
            checkedListBox1.Items.Add("Ponedjeljak", CheckState.Unchecked);
            checkedListBox1.Items.Add("Utorak", CheckState.Indeterminate);
            checkedListBox1.Items.Add("Srijeda", CheckState.Checked);
            checkedListBox1.Items.Add("Cetvrtak", CheckState.Unchecked);
            checkedListBox1.Items.Add("Petak", CheckState.Indeterminate);
            checkedListBox1.Items.Add("Subota", CheckState.Indeterminate);
        
        }


      
    }

        
}
