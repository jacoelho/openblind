package reviews

import (
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

const fixture = `<li class="empReview cf " id="empReview_45005756" data-brandviews="MODULE:n=reviews-reviewsPage-review:eid=43129:review_id=45005756" data-triggered-brandview="">
<div class="gdReview">
	<div class="d-flex justify-content-between">
		<div class="d-flex align-items-center"><time class="date subtle small" datetime="Sun Apr 04 2021 17:00:47 GMT+0100 (British Summer Time)">4 April 2021</time></div>
	</div>
	<div class="row mt">
		<div class="col-sm-1"><span class="sqLogo smSqLogo logoOverlay"><img alt="Tesla Logo" class="lazy lazy-loaded" data-original="https://media.glassdoor.com/sql/43129/tesla-squarelogo-1609189398200.png" data-original-2x="https://media.glassdoor.com/sqll/43129/tesla-squarelogo-1609189398200.png" data-retina-ok="true" src="https://media.glassdoor.com/sql/43129/tesla-squarelogo-1609189398200.png" title="" style="opacity: 1;"></span></div>
		<div class="col-sm-11 pl-sm-lg  mx-0">
			<div class="">
				<h2 class="h2 summary strong mb-xsm mt-0"><a href="/Reviews/Employee-Review-Tesla-RVW45005756.htm" class="reviewLink">"Great Company"</a></h2>
				<div class="mr-xsm d-lg-inline-block">
					<span class="gdStars gdRatings subRatings__SubRatingsStyles__gdStars">
						<div class=" v2__EIReviewsRatingsStylesV2__ratingInfoWrapper">
							<div class="v2__EIReviewsRatingsStylesV2__ratingInfo" rel="nofollow">
								<div class="v2__EIReviewsRatingsStylesV2__ratingNum v2__EIReviewsRatingsStylesV2__small">5.0</div>
								<span class="gdStars gdRatings common__StarStyles__gdStars">
									<span class="rating"><span title="5.0"></span></span>
									<div font-size="sm" class="css-1dc0bv4"><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span></div>
								</span>
								<span class="SVGInline">
									<svg class="SVGInline-svg" style="width: 16;height: 16;" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
										<path d="M4.4 9.25l7.386 7.523a1 1 0 001.428 0L20.6 9.25c.5-.509.5-1.324 0-1.833a1.261 1.261 0 00-1.8 0l-6.3 6.416-6.3-6.416a1.261 1.261 0 00-1.8 0c-.5.509-.5 1.324 0 1.833z" fill-rule="evenodd" fill="currentColor"></path>
									</svg>
								</span>
							</div>
						</div>
						<div class="subRatings module subRatings__SubRatingsStyles__subRatings">
							<div class="dummyHoverArea"></div>
							<i class="beak subRatings__SubRatingsStyles__beak"></i>
							<ul class="undecorated">
								<li>
									<div class="minor">Work/Life Balance</div>
									<span class="subRatings__SubRatingsStyles__gdBars gdBars gdRatings med" title="2.0">
										<span class="gdStars gdRatings common__StarStyles__gdStars">
											<span class="rating"><span title="2.0"></span></span>
											<div font-size="sm" class="css-19o85uz"><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span></div>
										</span>
									</span>
								</li>
								<li>
									<div class="minor">Culture &amp; Values</div>
									<span class="subRatings__SubRatingsStyles__gdBars gdBars gdRatings med" title="5.0">
										<span class="gdStars gdRatings common__StarStyles__gdStars">
											<span class="rating"><span title="5.0"></span></span>
											<div font-size="sm" class="css-1dc0bv4"><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span></div>
										</span>
									</span>
								</li>
								<li>
									<div class="minor">Diversity &amp; Inclusion</div>
									<span class="subRatings__SubRatingsStyles__gdBars gdBars gdRatings med" title="5.0">
										<span class="gdStars gdRatings common__StarStyles__gdStars">
											<span class="rating"><span title="5.0"></span></span>
											<div font-size="sm" class="css-1dc0bv4"><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span></div>
										</span>
									</span>
								</li>
								<li>
									<div class="minor">Career Opportunities</div>
									<span class="subRatings__SubRatingsStyles__gdBars gdBars gdRatings med" title="5.0">
										<span class="gdStars gdRatings common__StarStyles__gdStars">
											<span class="rating"><span title="5.0"></span></span>
											<div font-size="sm" class="css-1dc0bv4"><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span></div>
										</span>
									</span>
								</li>
								<li>
									<div class="minor">Compensation and Benefits</div>
									<span class="subRatings__SubRatingsStyles__gdBars gdBars gdRatings med" title="3.0">
										<span class="gdStars gdRatings common__StarStyles__gdStars">
											<span class="rating"><span title="3.0"></span></span>
											<div font-size="sm" class="css-1ihykkv"><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span></div>
										</span>
									</span>
								</li>
								<li>
									<div class="minor">Senior Management</div>
									<span class="subRatings__SubRatingsStyles__gdBars gdBars gdRatings med" title="4.0">
										<span class="gdStars gdRatings common__StarStyles__gdStars">
											<span class="rating"><span title="4.0"></span></span>
											<div font-size="sm" class="css-1c07csa"><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span><span class="gd-ui-star  css-fosmlm" role="button" color="#0caa41" font-size="sm" tabindex="0">★</span></div>
										</span>
									</span>
								</li>
							</ul>
						</div>
					</span>
				</div>
				<div class="d-lg-inline-block">
					<div class="author minor"><span class="authorInfo"><span class="authorJobTitle middle ">Current Employee - Global Supply Analyst</span>&nbsp;<span class="middle">in <span class="authorLocation">San Francisco, CA</span></span></span></div>
				</div>
				<div>
					<div class="row reviewBodyCell recommends">
						<div class="col-sm-4 d-flex align-items-center"><i class="sqLed middle sm mr-xsm green"></i><span>Recommends</span></div>
						<div class="col-sm-4 d-flex align-items-center"><i class="sqLed middle sm mr-xsm green"></i><span>Positive Outlook</span></div>
						<div class="col-sm-4 d-flex align-items-center"><i class="sqLed middle sm mr-xsm green"></i><span>Approves of CEO</span></div>
					</div>
				</div>
				<p class="mainText mb-0">I have been working at Tesla full-time</p>
			</div>
			<div class="">
				<div class="v2__EIReviewDetailsV2__fullWidth ">
					<p class="mb-0 strong mt-xsm">Pros</p>
					<p class="mt-0 mb-xsm v2__EIReviewDetailsV2__bodyColor v2__EIReviewDetailsV2__lineHeightLarge v2__EIReviewDetailsV2__isExpanded  "><span data-test="pros">Amazing work, very involved in day-to-day details of the company.</span></p>
				</div>
				<div class="v2__EIReviewDetailsV2__fullWidth ">
					<p class="mb-0 strong mt-xsm">Cons</p>
					<p class="mt-0 mb-xsm v2__EIReviewDetailsV2__bodyColor v2__EIReviewDetailsV2__lineHeightLarge v2__EIReviewDetailsV2__isExpanded  "><span data-test="cons">Work-life balance is not the best.</span></p>
				</div>
				<div class="row mt-xsm mx-0"></div>
				<div class="
					justify-content-around justify-content-md-between
					mt-lg row
					">
					<div class="shareContent d-flex justify-content-center">
						<div class="share-callout-inline">
							<div class="callout-container">
								<ul class="d-table social-share-icon-list p-0">
									<li class="cell"><a class="social-share-icon facebook-share" href="#shareOnFacebook" data-url="http://www.glassdoor.co.uk/Reviews/Employee-Review-Tesla-RVW45005756.htm" data-label="facebook" data-reviewid="45005756"><span class="offScreen">Share on Facebook</span></a></li>
									<li class="cell"><a class="social-share-icon twitter-share" href="https://twitter.com/share?url=http://www.glassdoor.co.uk/Reviews/Employee-Review-Tesla-RVW45005756.htm&amp;text=Tesla+review+on+%23Glassdoor%3A+%22Great Company%22" data-label="twitter" data-reviewid="45005756" rel="noopener noreferrer" target="_blank"><span class="offScreen">Share on Twitter</span></a></li>
									<li class="cell whatsapp"><a class="social-share-icon whatsapp-share" href="whatsapp://send?text=http://www.glassdoor.co.uk/Reviews/Employee-Review-Tesla-RVW45005756.htm" data-reviewid="45005756" data-action="share/whatsapp/share" data-label="whatsapp" rel="noopener noreferrer" target="_blank"><span class="offScreen">Share on WhatsApp</span></a></li>
									<li class="cell"><a class="social-share-icon email-share" href="mailto:?Subject=Tesla review on Glassdoor&amp;body=Read this review of Tesla on Glassdoor. %22Great Company%22&nbsp;http://www.glassdoor.co.uk/Reviews/Employee-Review-Tesla-RVW45005756.htm" data-label="email" data-reviewid="45005756" rel="noopener noreferrer" target="_blank"><span class="offScreen">Share via Email</span></a></li>
									<li class="cell"><a class="social-share-icon link-share" href="http://www.glassdoor.co.uk/Reviews/Employee-Review-Tesla-RVW45005756.htm" data-reviewid="45005756" data-label="link"><span class="offScreen">Copy Link</span></a></li>
									<li class="cell linkCopySuccess"><span class="social-share-icon icon-check showDesk"></span><span>Link Copied!</span></li>
								</ul>
							</div>
						</div>
					</div>
					<div class="d-flex">
						<div class="mr-md"><button class="gd-ui-button  css-glrvaa">Helpful </button></div>
						<div class=""><span class="flagContent" data-disp-type="review" data-id="45005756" data-member="true" data-review-link="/Reviews/Employee-Review-Tesla-RVW45005756.htm" data-type="EMPLOYER_REVIEW"><button class="px-0 mx-0 simple gd-btn gd-btn-2 gd-btn-sm gd-btn-icon gradient" title="Flag as Inappropriate" type="button"><i class="icon-flag-content "><span>Flag as Inappropriate</span></i><i class="hlpr"></i><span class="offScreen">Flag as Inappropriate</span></button><span class="posPt"></span></span></div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
</li>`

func TestParseInterview(t *testing.T) {
	root, err := html.Parse(strings.NewReader(fixture))
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	got, err := parseReview(root)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	want := Review{
		ID:     "45005756",
		Date:   mustParseTime(t, "2021-04-04T16:00:47Z"),
		Title:  `"Great Company"`,
		Rating: 5.0,
		Pros:   []string{"Amazing work, very involved in day-to-day details of the company."},
		Cons:   []string{"Work-life balance is not the best."},
		Advice: nil,
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parseReview() mismatch (-want +got):\n%s", diff)
	}
}

func mustParseTime(t *testing.T, s string) time.Time {
	t.Helper()

	parsed, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatal(err)
	}

	return parsed
}
