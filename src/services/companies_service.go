package services

import "github.com/nilerajput91/go-gin-docker-aws/src/models"

func GetAllCompanies() []models.Company {
	companies := []models.Company{
		{
			ID:          "1",
			Name:        "TechGenius",
			Industry:    "Technology",
			Slogan:      "Innovate for Tomorrow",
			WebsiteURL:  "https://techgenius.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Douala", Venue: "Tech Street"},
			Since:       "2010",
			CEO:         "John Doe",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/techgenius", Instagram: "https://instagram.com/techgenius"},
		},
		{
			ID:          "2",
			Name:        "AgroHarvest",
			Industry:    "Agriculture",
			Slogan:      "Cultivating Quality, Nurturing Life",
			WebsiteURL:  "https://agroharvest.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Bamenda", Venue: "Green Valley"},
			Since:       "2012",
			CEO:         "Alice Green",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/agroharvest", Instagram: "https://instagram.com/agroharvest"},
		},
		{
			ID:          "3",
			Name:        "EcoTech Solutions",
			Industry:    "Environmental Services",
			Slogan:      "Sustainable Solutions for a Green Tomorrow",
			WebsiteURL:  "https://ecotech.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Yaoundé", Venue: "EcoPark"},
			Since:       "2015",
			CEO:         "Emma Eco",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/ecotech", Instagram: "https://instagram.com/ecotech"},
		},
		{
			ID:          "4",
			Name:        "FoodDelight",
			Industry:    "Culinary",
			Slogan:      "Taste the Delight, Feel the Flavor",
			WebsiteURL:  "https://fooddelight.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Limbe", Venue: "Gourmet Avenue"},
			Since:       "2013",
			CEO:         "Chef Delia",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/fooddelight", Instagram: "https://instagram.com/fooddelight"},
		},
		{
			ID:          "5",
			Name:        "CamFashion",
			Industry:    "Fashion",
			Slogan:      "Where Style Meets Elegance",
			WebsiteURL:  "https://camfashion.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Buea", Venue: "Fashion Street"},
			Since:       "2014",
			CEO:         "Stella Styles",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/camfashion", Instagram: "https://instagram.com/camfashion"},
		},
		{
			ID:          "6",
			Name:        "GreenPower",
			Industry:    "Renewable Energy",
			Slogan:      "Empowering Lives, Energizing Futures",
			WebsiteURL:  "https://greenpower.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Kribi", Venue: "EcoEnergy Park"},
			Since:       "2011",
			CEO:         "James Green",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/greenpower", Instagram: "https://instagram.com/greenpower"},
		},
		{
			ID:          "7",
			Name:        "HealthHub",
			Industry:    "Healthcare",
			Slogan:      "Caring for a Healthier Tomorrow",
			WebsiteURL:  "https://healthhub.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Ebolowa", Venue: "Wellness Avenue"},
			Since:       "2016",
			CEO:         "Dr. Grace Healthy",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/healthhub", Instagram: "https://instagram.com/healthhub"},
		},
		{
			ID:          "8",
			Name:        "TourCam",
			Industry:    "Tourism",
			Slogan:      "Explore, Experience, Enjoy",
			WebsiteURL:  "https://tourcam.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Bafoussam", Venue: "Adventure Plaza"},
			Since:       "2018",
			CEO:         "Tom Traveler",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/tourcam", Instagram: "https://instagram.com/tourcam"},
		},
		{
			ID:          "9",
			Name:        "EduTech Solutions",
			Industry:    "Education Technology",
			Slogan:      "Transforming Education, Empowering Minds",
			WebsiteURL:  "https://edutech.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Dschang", Venue: "Knowledge Hub"},
			Since:       "2017",
			CEO:         "Professor Ed Tech",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/edutech", Instagram: "https://instagram.com/edutech"},
		},
		{
			ID:          "10",
			Name:        "ArtGallery",
			Industry:    "Arts and Culture",
			Slogan:      "Where Creativity Speaks",
			WebsiteURL:  "https://artgallery.cm",
			Headquater:  models.Office{Country: "Cameroon", Town: "Bamako", Venue: "Creative Corner"},
			Since:       "2016",
			CEO:         "Alice Artist",
			SocialMedia: models.SocialMedia{Facebook: "https://facebook.com/artgallery", Instagram: "https://instagram.com/artgallery"},
		},
	}

	return companies
}
