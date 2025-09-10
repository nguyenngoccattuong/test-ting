package Logic

import (
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

// Get Kols from the database based on the range of pageIndex and pageSize
func GetKolLogic(pageIndex, pageSize int64) ([]*DTO.KolDTO, int64, error) {
	var kols []Models.Kol
	var totalCount int64

	Initializers.DB.Model(&Models.Kol{}).Where("active = ? AND \"isRemove\" = ?", true, false).Count(&totalCount)

	offset := (pageIndex - 1) * pageSize
	err := Initializers.DB.Where("active = ? AND \"isRemove\" = ?", true, false).
		Offset(int(offset)).Limit(int(pageSize)).Find(&kols).Error

	if err != nil {
		return nil, 0, err
	}

	var kolDTOs []*DTO.KolDTO
	for _, kol := range kols {
		kolDTOs = append(kolDTOs, &DTO.KolDTO{
			KolID:                kol.KolID,
			UserProfileID:        kol.UserProfileID,
			Language:             kol.Language,
			Education:            kol.Education,
			ExpectedSalary:       kol.ExpectedSalary,
			ExpectedSalaryEnable: kol.ExpectedSalaryEnable,
			ChannelSettingTypeID: kol.ChannelSettingTypeID,
			IDFrontURL:           kol.IDFrontURL,
			IDBackURL:            kol.IDBackURL,
			PortraitURL:          kol.PortraitURL,
			RewardID:             kol.RewardID,
			PaymentMethodID:      kol.PaymentMethodID,
			TestimonialsID:       kol.TestimonialsID,
			VerificationStatus:   kol.VerificationStatus,
			Enabled:              kol.Enabled,
			ActiveDate:           kol.ActiveDate,
			Active:               kol.Active,
			CreatedBy:            kol.CreatedBy,
			CreatedDate:          kol.CreatedDate,
			ModifiedBy:           kol.ModifiedBy,
			ModifiedDate:         kol.ModifiedDate,
			IsRemove:             kol.IsRemove,
			IsOnBoarding:         kol.IsOnBoarding,
			Code:                 kol.Code,
			PortraitRightURL:     kol.PortraitRightURL,
			PortraitLeftURL:      kol.PortraitLeftURL,
			LivenessStatus:       kol.LivenessStatus,
		})
	}

	return kolDTOs, totalCount, nil
}
