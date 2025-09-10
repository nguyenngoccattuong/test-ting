"use client"

import React, {useEffect, useState} from 'react';
import './page.css'

interface Kols {
	kolID: number;
	userProfileID: number;
	language: string;
	education: string;
	expectedSalary: number;
	expectedSalaryEnable: boolean;
	channelSettingTypeID: number;
	iDFrontURL: string;
	iDBackURL: string;
	portraitURL: string;
	rewardID: number;
	paymentMethodID: number;
	testimonialsID: number;
	verificationStatus: string;
	enabled: boolean;
	activeDate: string;
	active: boolean;
	createdBy: string;
	createdDate: string;
	modifiedBy: string;
	modifiedDate: string;
	isRemove: boolean;
	isOnBoarding: boolean;
	code: string;
	portraitRightURL: string;
	portraitLeftURL: string;
	livenessStatus: string;
}

interface ApiResponse {
	result: string;
	errorMessage: string;
	pageIndex: number;
	pageSize: number;
	totalCount: number;
	kolInformation: Kols[];
}

export default function Page() {
    const [Kols, setKols] = useState<Kols[]>([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);
	const [currentIndex, setCurrentIndex] = useState(0);
	const itemsPerPage = 3;

    useEffect(() => {
		const fetchKols = async () => {
			try {
				setLoading(true);
				const response = await fetch('http://localhost:8081/kols?pageIndex=1&pageSize=100');
				const data: ApiResponse = await response.json();
				
				if (data.result === 'Success') {
					setKols(data.kolInformation);
				} else {
					setError(data.errorMessage);
				}
			} catch (err) {
				setError('Failed to fetch KOLs');
				console.error('Error fetching KOLs:', err);
			} finally {
				setLoading(false);
			}
		};

		fetchKols();
    }, []);

	const scrollLeft = () => {
		setCurrentIndex(Math.max(0, currentIndex - 1));
	};

	const scrollRight = () => {
		setCurrentIndex(Math.min(Kols.length - itemsPerPage, currentIndex + 1));
	};

	const formatDate = (dateString: string) => {
		return new Date(dateString).toLocaleDateString('en-US', {
			year: 'numeric',
			month: 'short',
			day: 'numeric'
		});
	};

	const formatSalary = (salary: number) => {
		return new Intl.NumberFormat('en-US', {
			style: 'currency',
			currency: 'USD',
			minimumFractionDigits: 0
		}).format(salary);
	};

	if (loading) {
		return (
			<div className="container">
				<h1 className="header">KOL Management System</h1>
				<div className="loading-container">
					<div className="loading-spinner"></div>
					<h2 style={{ color: '#4a5568', marginBottom: '10px' }}>Loading KOLs...</h2>
					<p style={{ color: '#718096' }}>Please wait while we fetch the data</p>
				</div>
			</div>
		);
	}

	if (error) {
		return (
			<div className="container">
				<h1 className="header">KOL Management System</h1>
				<div className="error-container">
					<h2 style={{ color: '#e53e3e', marginBottom: '20px' }}>⚠️ Error Loading Data</h2>
					<p style={{ color: '#4a5568', marginBottom: '20px' }}>{error}</p>
					<button 
						onClick={() => window.location.reload()} 
						style={{
							background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
							color: 'white',
							border: 'none',
							padding: '12px 24px',
							borderRadius: '12px',
							cursor: 'pointer',
							fontWeight: '600'
						}}
					>
						🔄 Retry
					</button>
				</div>
			</div>
		);
	}

    return (
        <div className="container">
            <h1 className="header">KOL Management System</h1>
			<div className="subtitle">Total KOLs: {Kols.length}</div>
			
			<div className="kol-container">
				<button 
					className="scroll-btn left" 
					onClick={scrollLeft} 
					disabled={currentIndex === 0}
				>
					← Previous
				</button>
				
				<div className="kol-list">
					{Kols.slice(currentIndex, currentIndex + itemsPerPage).map(kol => (
						<div key={kol.kolID} className="kol-card">
							<div className="kol-header">
								<h3 className="kol-code">{kol.code}</h3>
								<span className={`status-badge ${kol.verificationStatus.toLowerCase()}`}>
									{kol.verificationStatus}
								</span>
							</div>
							
							<div className="kol-info">
								<div className="info-row">
									<span className="label">🌐 Language:</span>
									<span className="value">{kol.language.toUpperCase()}</span>
								</div>
								<div className="info-row">
									<span className="label">🎓 Education:</span>
									<span className="value">{kol.education}</span>
								</div>
								<div className="info-row">
									<span className="label">💰 Expected Salary:</span>
									<span className="value">
										{kol.expectedSalaryEnable ? formatSalary(kol.expectedSalary) : 'Not disclosed'}
									</span>
								</div>
								<div className="info-row">
									<span className="label">👁 Liveness Status:</span>
									<span className={`liveness-badge ${kol.livenessStatus.toLowerCase()}`}>
										{kol.livenessStatus}
									</span>
								</div>
								<div className="info-row">
									<span className="label">✅ Onboarding:</span>
									<span className={`onboarding-badge ${kol.isOnBoarding ? 'completed' : 'pending'}`}>
										{kol.isOnBoarding ? 'Completed' : 'Pending'}
									</span>
								</div>
								<div className="info-row">
									<span className="label">📅 Active Date:</span>
									<span className="value">{formatDate(kol.activeDate)}</span>
								</div>
							</div>
						</div>
					))}
				</div>
				
				<button 
					className="scroll-btn right" 
					onClick={scrollRight} 
					disabled={currentIndex >= Kols.length - itemsPerPage}
				>
					Next →
				</button>
			</div>
			
			<div className="pagination-info">
				Showing {currentIndex + 1}-{Math.min(currentIndex + itemsPerPage, Kols.length)} of {Kols.length} KOLs
			</div>
        </div>
    )
}