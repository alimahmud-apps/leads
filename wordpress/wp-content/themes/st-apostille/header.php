<!DOCTYPE html>
<html <?php language_attributes(); ?>>
<head>
    <meta charset="<?php bloginfo( 'charset' ); ?>">
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="profile" href="https://gmpg.org/xfn/11">
    <?php wp_head(); ?>
</head>

<body <?php body_class(); ?>>
    <?php
    if ( function_exists( 'wp_body_open' ) ) {
        wp_body_open();
    } else {
        do_action( 'wp_body_open' );
    }
    ?>
    <!-- Page Wrapper -->
    <div id="page-wrap">

        <a class="skip-link screen-reader-text" href="#skip-link-target"><?php esc_html_e( 'Skip to content', 'st-apostille' ); ?></a>

        <header id="masthead" class="site-header header py-4 sticky-header" role="banner" style="<?php header_textcolor(); ?>">
            <div class="container">

                <!-- Display Custom Header Image -->
                <?php if ( get_header_image() ) : ?>
                    <div id="custom-header" class="mb-4">
                        <img src="<?php echo esc_url( get_header_image() ); ?>" 
                             width="<?php echo esc_attr( get_custom_header()->width ); ?>" 
                             height="<?php echo esc_attr( get_custom_header()->height ); ?>" 
                             alt="<?php esc_attr_e( 'Header Image', 'st-apostille' ); ?>" />
                    </div>
                <?php endif; ?>

                <div class="row">
                    <!-- Logo and Site Identity -->
                    <div class="col-lg-3 col-md-3 align-self-center text-center site-logo">
                        <div class="logo text-center text-md-left mb-3 mb-md-0">
                            <?php
                            if ( has_custom_logo() ) :
                                the_custom_logo();
                            else :
                            ?>
                                <h1 class="site-title">
                                    <a href="<?php echo esc_url( home_url( '/' ) ); ?>" 
                                       title="<?php esc_attr_e( 'Home', 'st-apostille' ); ?>" 
                                       rel="home" style="color: #<?php echo esc_attr( get_header_textcolor() ); ?>;">
                                        <?php echo esc_html( get_bloginfo( 'name' ) ); ?>
                                    </a>
                                </h1>
                                <?php if ( get_bloginfo( 'description', 'display' ) ) : ?>
                                <p class="site-description" style="color: #<?php echo esc_attr( get_header_textcolor() ); ?>;">
                                    <?php echo esc_html( get_bloginfo( 'description', 'display' ) ); ?>
                                </p>
                                <?php endif; ?>
                            <?php endif; ?>
                        </div>
                    </div>

                    <!-- Navigation Menu -->
                        <nav id="site-navigation" class="main-navigation">
                            <button type="button" class="menu-toggle">
                                <i class="fa fa-list-ul fa-2x"></i>
                            </button>
                            <?php
                            wp_nav_menu(
                                array(
                                    'theme_location' => 'primary',
                                    'menu_id'        => 'primary-menu',
                                )
                            );
                            ?>
                        </nav><!-- #site-navigation -->
                </div>
            </div> 
        </header>
    </div> <!-- End of Page Wrapper -->
