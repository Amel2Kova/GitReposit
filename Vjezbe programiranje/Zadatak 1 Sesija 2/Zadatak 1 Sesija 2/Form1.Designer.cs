namespace zadatak1
{
    partial class Form1
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.Dorucak = new System.Windows.Forms.CheckBox();
            this.Rucak = new System.Windows.Forms.CheckBox();
            this.button1 = new System.Windows.Forms.Button();
            this.Vecera = new System.Windows.Forms.CheckBox();
            this.checkBox4 = new System.Windows.Forms.CheckBox();
            this.textBox1 = new System.Windows.Forms.TextBox();
            this.SuspendLayout();
            // 
            // Dorucak
            // 
            this.Dorucak.AutoSize = true;
            this.Dorucak.Location = new System.Drawing.Point(69, 62);
            this.Dorucak.Name = "Dorucak";
            this.Dorucak.Size = new System.Drawing.Size(67, 17);
            this.Dorucak.TabIndex = 0;
            this.Dorucak.Text = "Dorucak";
            this.Dorucak.UseVisualStyleBackColor = true;
            // 
            // Rucak
            // 
            this.Rucak.AutoSize = true;
            this.Rucak.Location = new System.Drawing.Point(69, 86);
            this.Rucak.Name = "Rucak";
            this.Rucak.Size = new System.Drawing.Size(58, 17);
            this.Rucak.TabIndex = 1;
            this.Rucak.Text = "Rucak";
            this.Rucak.UseVisualStyleBackColor = true;
            // 
            // button1
            // 
            this.button1.Location = new System.Drawing.Point(69, 156);
            this.button1.Name = "button1";
            this.button1.Size = new System.Drawing.Size(75, 23);
            this.button1.TabIndex = 2;
            this.button1.Text = "Posalji";
            this.button1.UseVisualStyleBackColor = true;
            this.button1.Click += new System.EventHandler(this.button1_Click);
            // 
            // Vecera
            // 
            this.Vecera.AutoSize = true;
            this.Vecera.Location = new System.Drawing.Point(69, 110);
            this.Vecera.Name = "Vecera";
            this.Vecera.Size = new System.Drawing.Size(60, 17);
            this.Vecera.TabIndex = 3;
            this.Vecera.Text = "Vecera";
            this.Vecera.UseVisualStyleBackColor = true;
            // 
            // checkBox4
            // 
            this.checkBox4.AutoSize = true;
            this.checkBox4.Location = new System.Drawing.Point(171, 161);
            this.checkBox4.Name = "checkBox4";
            this.checkBox4.Size = new System.Drawing.Size(107, 17);
            this.checkBox4.TabIndex = 4;
            this.checkBox4.Text = "Sakrij porudzbinu";
            this.checkBox4.UseVisualStyleBackColor = true;
            this.checkBox4.CheckedChanged += new System.EventHandler(this.checkBox4_CheckedChanged);
            // 
            // textBox1
            // 
            this.textBox1.Location = new System.Drawing.Point(171, 62);
            this.textBox1.Multiline = true;
            this.textBox1.Name = "textBox1";
            this.textBox1.Size = new System.Drawing.Size(153, 65);
            this.textBox1.TabIndex = 5;
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(418, 259);
            this.Controls.Add(this.textBox1);
            this.Controls.Add(this.checkBox4);
            this.Controls.Add(this.Vecera);
            this.Controls.Add(this.button1);
            this.Controls.Add(this.Rucak);
            this.Controls.Add(this.Dorucak);
            this.Name = "Form1";
            this.Text = "Form1";
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.CheckBox Dorucak;
        private System.Windows.Forms.CheckBox Rucak;
        private System.Windows.Forms.Button button1;
        private System.Windows.Forms.CheckBox Vecera;
        private System.Windows.Forms.CheckBox checkBox4;
        private System.Windows.Forms.TextBox textBox1;
    }
}

