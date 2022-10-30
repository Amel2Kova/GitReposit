using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Projekt
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            foreach (Control x in this.Controls) {
                if (x.Tag == "Platform") { x.BringToFront(); }
                if (x.Tag == "Coin") { x.SendToBack(); }

            }


        }

        private void player_Click(object sender, EventArgs e)
        {

        }

        private void pictureBox21_Click(object sender, EventArgs e)
        {

        }
        bool GoLeft = false, GoRight = false, jumping = false, isGameOver = false;
        int jumpSpeed = 0;
        int GRAVITY = 9;
        int force = 8;
        int score = 0;
        int playerspeed = 7;
        int horisontalSpeed = 5;
        int verticalSpeed = 3;
        int enemySpeed = 4;
        int Speed = 5;

        private void KeyIsDown(object sender, KeyEventArgs e)
        {
            if (e.KeyCode == Keys.Left)
            { GoLeft = true; }

            if (e.KeyCode == Keys.Right)
            { GoRight = true; }

            if (e.KeyCode == Keys.Up && jumping == false)
            { jumping = true; }
        }

        private void KeyIsUo(object sender, KeyEventArgs e)
        {
            if (e.KeyCode == Keys.Left)
            { GoLeft = false; }

            if (e.KeyCode == Keys.Right)
            { GoRight = false; }

            if (jumping == true) { jumping = false; }

            if (e.KeyCode == Keys.Space && isGameOver == true) { Restart(); }
        }

        private void MAinGameTic(object sender, EventArgs e)
        {
            label1.Text = "Score is : " + score;

            player.Top += (GRAVITY + jumpSpeed);

            if (GoLeft == true) { player.Left -= playerspeed; }
            if (GoRight == true) { player.Left += playerspeed; }
            if (jumping == true && force < 0) { jumping = false; }
            if (jumping == true) { jumpSpeed = -20; force -= 1; GRAVITY = 10; }
            else
            {
                jumpSpeed = 0; GRAVITY = 10;

                foreach (Control x in this.Controls)
                {
                    if (x is PictureBox)
                    {
                        if (x.Tag == "PlatformDown")
                        {
                            if (player.Bounds.IntersectsWith(x.Bounds)) { player.Top = x.Top + player.Height+x.Height;
                            force = 0;
                            
                            }
                            
                        }
                        if (x.Tag == "Platform")
                        {
                            if (player.Bounds.IntersectsWith(x.Bounds))
                            {
                                force = 8;
                                GRAVITY = 0;
                                 player.Top = x.Top - player.Height;
                            }
                            
                        }
                        if (x.Tag == "coin")
                        {
                            if (player.Bounds.IntersectsWith(x.Bounds) && x.Visible == true)
                            {
                                score += 1;
                                x.Visible = false;
                                label1.Text = "Score is : " + score;
                            }
                        }
                        if (x.Tag == "enemy")
                        {
                            if (player.Bounds.IntersectsWith(x.Bounds))
                            {
                                timer1.Stop();
                                isGameOver = true;
                                label1.Text = "Score is : " + score + Environment.NewLine + "GAME OWER";
                            }
                        }
                        if (x.Tag == "Vrata")
                        {
                            if (player.Bounds.IntersectsWith(x.Bounds))
                            {
                                timer1.Stop();
                                isGameOver = true;
                                label1.Text = "Score is : " + score + Environment.NewLine + "YOU WIN!!!";
                            }
                        }



                    }

                }
                
                HorisontalPlatform.Left += horisontalSpeed;
                if (HorisontalPlatform.Left < 0 || HorisontalPlatform.Left + HorisontalPlatform.Width > this.ClientSize.Width)
                { horisontalSpeed = -horisontalSpeed; }
                VerticalPlatform.Top -= verticalSpeed;
                if (VerticalPlatform.Top < 0 || VerticalPlatform.Top > this.ClientSize.Height - VerticalPlatform.Height)
                { verticalSpeed = -verticalSpeed; }
                enemy1.Left += Speed;
                if (enemy1.Left < pictureBox3.Left || (enemy1.Left + enemy1.Width) >( pictureBox3.Left + pictureBox3.Width))
                {
                    Speed = -Speed;

                }
                enemy2.Left += enemySpeed;
                if (enemy2.Left < pictureBox7.Left || enemy2.Left + enemy2.Width > pictureBox7.Left + pictureBox7.Width)
                {
                    enemySpeed = -enemySpeed;

                }
                if (player.Top > (this.ClientSize.Height + player.Height))
                {
                    timer1.Stop();
                    isGameOver = true;
                    label1.Text = "Score is : " + score + Environment.NewLine + "GAME OWER";
                }
            }
        }

            private void Restart() {
                jumping = false;
                GoLeft = false;
                GoRight = false;
                isGameOver = false;
                score = 0;
                label1.Text = "Score is : " + score;
                foreach (Control item in this.Controls)
                {// Postavimo da su svi coini vidljivi
                    if (item is PictureBox && item.Visible == false) {
                        item.Visible = true;
                    }
                }
                //Restartujemo poziciju igraca, protivnika i kretajucih platformi

                player.Left = 24;
                player.Top = 348;
                enemy1.Left = 505;
                enemy2.Left = 321;
                VerticalPlatform.Top = 237;
                HorisontalPlatform.Left = 58;
                timer1.Start();
            }
        }
    }
