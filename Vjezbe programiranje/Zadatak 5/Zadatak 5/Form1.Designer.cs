namespace Zadatak_5
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
            this.labelPrvi = new System.Windows.Forms.Label();
            this.labelDrugi = new System.Windows.Forms.Label();
            this.txtPrvi = new System.Windows.Forms.TextBox();
            this.txtDrugi = new System.Windows.Forms.TextBox();
            this.cmdIzracunaj = new System.Windows.Forms.Button();
            this.SuspendLayout();
            // 
            // labelPrvi
            // 
            this.labelPrvi.AutoSize = true;
            this.labelPrvi.Location = new System.Drawing.Point(25, 36);
            this.labelPrvi.Name = "labelPrvi";
            this.labelPrvi.Size = new System.Drawing.Size(45, 13);
            this.labelPrvi.TabIndex = 0;
            this.labelPrvi.Text = "Prvi broj";
            // 
            // labelDrugi
            // 
            this.labelDrugi.AutoSize = true;
            this.labelDrugi.Location = new System.Drawing.Point(25, 70);
            this.labelDrugi.Name = "labelDrugi";
            this.labelDrugi.Size = new System.Drawing.Size(52, 13);
            this.labelDrugi.TabIndex = 1;
            this.labelDrugi.Text = "Drugi broj";
            // 
            // txtPrvi
            // 
            this.txtPrvi.Location = new System.Drawing.Point(100, 29);
            this.txtPrvi.Name = "txtPrvi";
            this.txtPrvi.Size = new System.Drawing.Size(100, 20);
            this.txtPrvi.TabIndex = 2;
            this.txtPrvi.TextChanged += new System.EventHandler(this.txtPrvi_TextChanged);
            // 
            // txtDrugi
            // 
            this.txtDrugi.Location = new System.Drawing.Point(100, 63);
            this.txtDrugi.Name = "txtDrugi";
            this.txtDrugi.Size = new System.Drawing.Size(100, 20);
            this.txtDrugi.TabIndex = 3;
            // 
            // cmdIzracunaj
            // 
            this.cmdIzracunaj.Location = new System.Drawing.Point(90, 122);
            this.cmdIzracunaj.Name = "cmdIzracunaj";
            this.cmdIzracunaj.Size = new System.Drawing.Size(75, 23);
            this.cmdIzracunaj.TabIndex = 4;
            this.cmdIzracunaj.Text = "&Izračunaj zbir";
            this.cmdIzracunaj.UseVisualStyleBackColor = true;
            this.cmdIzracunaj.Click += new System.EventHandler(this.cmdIzracunaj_Click);
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(258, 155);
            this.Controls.Add(this.cmdIzracunaj);
            this.Controls.Add(this.txtDrugi);
            this.Controls.Add(this.txtPrvi);
            this.Controls.Add(this.labelDrugi);
            this.Controls.Add(this.labelPrvi);
            this.Name = "Form1";
            this.Text = "Obrada izuzetaka";
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.Label labelPrvi;
        private System.Windows.Forms.Label labelDrugi;
        private System.Windows.Forms.TextBox txtPrvi;
        private System.Windows.Forms.TextBox txtDrugi;
        private System.Windows.Forms.Button cmdIzracunaj;
    }
}

