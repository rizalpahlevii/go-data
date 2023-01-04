package main

import (
	"fmt"
	"main/configurations"
)

//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `customer12692`
//--
//
//CREATE TABLE `customer12692` (
//`customer_id` varchar(255) NOT NULL,
//`customer_unique_id` varchar(255) NOT NULL,
//`customer_zip_code_prefix` varchar(255) NOT NULL,
//`customer_city` varchar(255) NOT NULL,
//`customer_state` varchar(255) NOT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `geolocation12692`
//--
//
//CREATE TABLE `geolocation12692` (
//`geolocation_zip_code_prefix` varchar(50) DEFAULT NULL,
//`geolocation_lat` varchar(50) DEFAULT NULL,
//`geolocation_lng` varchar(50) DEFAULT NULL,
//`geolocation_city` varchar(50) DEFAULT NULL,
//`geolocation_state` varchar(50) DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `location12692`
//--
//
//CREATE TABLE `location12692` (
//`state` varchar(255) DEFAULT NULL,
//`state_code` varchar(255) DEFAULT NULL,
//`city` varchar(255) DEFAULT NULL,
//`zip_code_start` varchar(255) DEFAULT NULL,
//`zip_code_end` varchar(255) DEFAULT NULL,
//`state_lat` varchar(255) DEFAULT NULL,
//`state_lng` varchar(255) DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `order12692`
//--
//
//CREATE TABLE `order12692` (
//`order_id` varchar(50) NOT NULL,
//`customer_id` varchar(50) NOT NULL,
//`order_status` varchar(15) NOT NULL,
//`order_purchase_timestamp` timestamp NOT NULL,
//`order_approved_at` timestamp NULL DEFAULT NULL,
//`order_delivered_carrier_date` timestamp NULL DEFAULT NULL,
//`order_delivered_customer_date` timestamp NULL DEFAULT NULL,
//`order_estimated_delivery_date` timestamp NOT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `order_item12692`
//--
//
//CREATE TABLE `order_item12692` (
//`order_id` varchar(50) NOT NULL,
//`order_item_id` varchar(50) NOT NULL,
//`product_id` varchar(50) NOT NULL,
//`seller_id` varchar(50) NOT NULL,
//`shipping_limit_date` timestamp NOT NULL,
//`price` float NOT NULL,
//`freight_value` float NOT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `order_payment12692`
//--
//
//CREATE TABLE `order_payment12692` (
//`order_id` varchar(50) DEFAULT NULL,
//`payment_sequential` int DEFAULT NULL,
//`payment_type` varchar(50) DEFAULT NULL,
//`payment_installments` int DEFAULT NULL,
//`payment_value` float DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `order_review12692`
//--
//
//CREATE TABLE `order_review12692` (
//`review_id` varchar(50) DEFAULT NULL,
//`order_id` varchar(50) DEFAULT NULL,
//`review_score` int DEFAULT NULL,
//`review_comment_title` varchar(255) DEFAULT NULL,
//`review_comment_message` varchar(255) DEFAULT NULL,
//`review_creation_date` datetime DEFAULT NULL,
//`review_answer_timestamp` timestamp NULL DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `product12692`
//--
//
//CREATE TABLE `product12692` (
//`product_id` varchar(255) NOT NULL,
//`product_category_name` varchar(255) DEFAULT NULL,
//`product_name_length` int DEFAULT NULL,
//`product_description_length` int DEFAULT NULL,
//`product_photos_qty` int DEFAULT NULL,
//`product_weight_g` int DEFAULT NULL,
//`product_length_cm` int DEFAULT NULL,
//`product_height_cm` int DEFAULT NULL,
//`product_width_cm` int DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `seller12692`
//--
//
//CREATE TABLE `seller12692` (
//`seller_id` varchar(50) NOT NULL,
//`seller_zip_code_prefix` varchar(7) NOT NULL,
//`seller_city` varchar(255) DEFAULT NULL,
//`seller_state` varchar(255) DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//-- --------------------------------------------------------
//
//--
//-- Table structure for table `translation12692`
//--
//
//CREATE TABLE `translation12692` (
//`product_category_name` varchar(255) NOT NULL,
//`product_category_name_english` varchar(255) NOT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
//
//--
//-- Indexes for dumped tables
//--
//
//--
//-- Indexes for table `customer12692`
//--
//ALTER TABLE `customer12692`
//ADD PRIMARY KEY (`customer_id`);
//
//--
//-- Indexes for table `order12692`
//--
//ALTER TABLE `order12692`
//ADD PRIMARY KEY (`order_id`);
//
//--
//-- Indexes for table `product12692`
//--
//ALTER TABLE `product12692`
//ADD PRIMARY KEY (`product_id`);
//
//--
//-- Indexes for table `seller12692`
//--
//ALTER TABLE `seller12692`
//ADD PRIMARY KEY (`seller_id`);
//COMMIT;
//
///*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
///*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
///*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

func getRevenues() {
	locationTable := configurations.Configuration().CustomerTableName
	geolocationTable := configurations.Configuration().GeolocationTableName
	orderTable := configurations.Configuration().OrderTableName
	orderItemTable := configurations.Configuration().OrderItemTableName
	productTable := configurations.Configuration().ProductTableName
	translationTable := configurations.Configuration().TranslationTableName
	orderPaymentTable := configurations.Configuration().OrderPaymentTableName
	orderReviewTable := configurations.Configuration().OrderReviewTableName
	sellerTable := configurations.Configuration().SellerTableName
	customerTable := configurations.Configuration().CustomerTableName

	//db := database.GetConnection()
	//defer db.Close()
	fmt.Println("TestGetRevenues")
	fmt.Println(locationTable, geolocationTable, orderTable, orderItemTable, productTable, translationTable, orderPaymentTable, orderReviewTable, sellerTable, customerTable)

}
