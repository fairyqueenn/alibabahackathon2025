<meta charset="UTF-8">
<meta content="IE=edge" http-equiv="X-UA-Compatible">
<?php if (@$title === NULL): ?>
    <title>App</title>
<?php else: ?>
    <title> <?= $title; ?> - App</title>
<?php endif; ?>
<meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
<meta content="Muhammad Saleh Solahudin, m.saleh.solahudin@gmail.com" name="author">
<meta content="Muhammad Saleh Solahudin" name="owner">
<?php if (@$description === NULL): ?>
  <?php if (@$no_description === NULL): ?>
    <meta content="?" name="description">
  <?php endif; ?>
<?php else: ?>
  <meta content="<?= strip_tags($description); ?>" name="description">
<?php endif; ?>
<meta content="App - ?" name="copyright">
<meta content="worldwide" name="coverage">
<meta content="global" name="distribution">
<meta content="general" name="rating">
<meta content="1 days" name="revisit-after">
<meta content="<?= $url; ?>" name="url">
<meta content="App - ?" name="application-name">
<!-- <meta content="#DC3545" name="theme-color">
<meta content="#DC3545" name="msapplication-navbutton-color">
<meta content="#DC3545" name="msapplication-TileColor">
<meta content="#DC3545" name="apple-mobile-web-app-status-bar-style"> -->
<meta content="App - ?" name="apple-mobile-web-app-title">
<meta content="yes" name="apple-touch-fullscreen">
<meta content="yes" name="apple-mobile-web-app-capable">
<meta content="<?= $url; ?>" name="msapplication-starturl">
<meta content="true" name="mssmarttagspreventparsing">
<meta content="all" property="webcrawlers">
<meta content="all" property="spiders">
<meta content="all" property="robots">
<meta content="id" name="language">
<meta content="id" name="geo.country">
<meta content="ID-JB" name="geo.region">
<meta content="Indonesia" name="geo.placename">
<meta content="-0.789275; 113.921327" name="geo.position">
<meta content="-0.789275, 113.921327" name="ICBM">
<meta content="summary_large_image" name="twitter:card">
<meta content="@msalehsolahudin" name="twitter:site">
<meta content="@msalehsolahudin" name="twitter:creator">
<meta content="<?= $url; ?>" name="twitter:url">
<meta content="website" property="og:type">
<meta content="<?= $url; ?>" property="og:url">
<?php if (@$sosmed_image === NULL): ?>
  <!-- <meta content="<?= base_url('media/website/banner.png'); ?>" name="twitter:image:src">
  <meta content="<?= base_url('media/website/banner.png'); ?>" property="og:image"> -->
<?php else: ?>
  <!-- <meta content="<?= $sosmed_image; ?>" name="twitter:image:src">
  <meta content="<?= $sosmed_image; ?>" property="og:image"> -->
<?php endif; ?>
<?php if (@$sosmed_meta_title === NULL): ?>
  <meta content="App - ?" name="twitter:title">
  <meta content="App - ?" property="og:site_name">
  <meta content="App - ?" property="og:title">
  <meta content="?" name="twitter:description">
  <meta content="?" property="og:description">
<?php else: ?>
  <meta content="<?= $sosmed_meta_title; ?>" name="twitter:title">
  <meta content="<?= $sosmed_meta_title; ?>" property="og:site_name">
  <meta content="<?= $sosmed_meta_title; ?>" property="og:title">
  <meta content="<?= $sosmed_meta_desc; ?>" name="twitter:description">
  <meta content="<?= $sosmed_meta_desc; ?>" property="og:description">
<?php endif; ?>
<meta content="id_ID" property="og:locale">
