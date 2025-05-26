<?php
defined('BASEPATH') or exit('No direct script access allowed');

class App extends CI_Controller
{
    public function __construct()
    {
        parent::__construct();
    }

    public function index()
    {
        $this->load->view('home', NULL);
    }

    public function add()
    {
        $this->load->view('add', NULL);
    }

    public function detail($id)
    {
        $this->load->view('detail', ["id" => $id]);
    }

    public function order()
    {
        $this->load->view('order', NULL);
    }
}
