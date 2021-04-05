package interviews

import (
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

const fixture = `<div class="mt-0 mb-0 my-md-std css-1aiqpnt pb-std" data-test="Interview44944117Container" data-brandviews="MODULE:n=ei-interviews-interview:eid=43129:interview_id=44944117" data-triggered-brandview="">
<div class="css-1qmwqit mb-md-md mb-xsm d-flex justify-content-between" data-test="Interview44944117Topline"><time datetime="2021-4-2">2 Apr 2021</time></div>
<div class="row">
	<div class="d-none d-md-block col-md-1" data-test="Interview44944117EmployerLogo"><span class="d-inline-block css-nk3mpl e151mjlk2"><img class="css-187fu8i" src="https://media.glassdoor.com/sql/43129/tesla-squarelogo-1609189398200.png" alt="Tesla" width="70"></span></div>
	<div class="col-12 col-md-11 pl-md-lg" data-test="Interview44944117Details">
		<div class="d-flex align-items-center">
			<h2 class="mt-0 strong mb-xsm" data-test="Interview44944117Title"><a href="/Interview/Tesla-Interview-RVW44944117.htm">Mechanical Engineer Intern Interview</a></h2>
		</div>
		<div class="mb-md css-1yyfea9 e151mjlk0" data-test="Interview44944117CandidateSubtext">Anonymous Employee</div>
		<div class="row">
			<div class="col-12 col-md-4 d-flex align-items-center mb-std" data-test="Interview44944117Rating"><span class="d-inline-block mr-xxsm green css-ozq8ud e11p9wri0"></span>Accepted Offer</div>
			<div class="col-12 col-md-4 d-flex align-items-center mb-std" data-test="Interview44944117Rating"><span class="d-inline-block mr-xxsm green css-ozq8ud e11p9wri0"></span>Positive Experience</div>
			<div class="col-12 col-md-4 d-flex align-items-center mb-std" data-test="Interview44944117Rating"><span class="d-inline-block mr-xxsm yellow css-ozq8ud e11p9wri0"></span>Average Interview</div>
		</div>
		<div>
			<div class="mt-sm" data-test="Interview44944117ApplicationDetails">
				<strong class="d-block">Application</strong>
				<p class="mt-xsm mb-std">I interviewed at Tesla</p>
			</div>
			<strong>Interview</strong>
			<p class="css-lyyc14 css-w00cnv  mb-std" data-test="Interview44944117Process">Highly flexible depending on team and directly interviewed by the team member, so could be just book technician questions or design scenarios. The number of times you get interviewed is also dependent on the team.</p>
			<button class="strong mb-std css-1e8g7ps eorog470">Continue Reading</button>
			<div data-test="Interview44944117QuestionsContainer">
				<strong class="d-block mb-xsm">Interview Questions</strong>
				<ul class="css-w00cnv pl-0 css-o9b79t e151mjlk3" data-test="Interview44944117Questions">
					<li class="mb-std">
						<span class="d-inline-block mb-sm">Why do you want to work for Tesla?</span>
						<div><a class=" css-1nx24df e151mjlk1" href="/Interview/Why-do-you-want-to-work-for-Tesla-QTN_4358096.htm">Answer Question</a></div>
					</li>
				</ul>
			</div>
		</div>
	</div>
</div>
<div class="d-flex flex-column flex-md-row align-items-start align-items-md-start justify-content-between" data-test="Interview44944117BottomBar">
	<div class="shareContent d-flex justify-content-center">
		<div class="share-callout-inline">
			<div class="callout-container">
				<ul class="d-table social-share-icon-list p-0" data-test="Interview44944117SocialButtons">
					<li class="cell middle"><a class="social-share-icon facebook-share" href="#shareOnFacebook" data-url="http://www.glassdoor.co.uk/Interview/Tesla-Interview-RVW44944117.htm" data-label="facebook" data-reviewid="44944117"><span class="offScreen">Share on Facebook</span></a></li>
					<li class="cell middle"><a class="social-share-icon twitter-share" href="https://twitter.com/share?url=http://www.glassdoor.co.uk/Interview/Tesla-Interview-RVW44944117.htm&amp;text=Tesla review on Glassdoor%22Mechanical Engineer Intern Interview%22" data-label="twitter" data-reviewid="44944117" rel="noopener noreferrer" target="_blank"><span class="offScreen">Share on Twitter</span></a></li>
					<li class="cell middle whatsapp"><a class="social-share-icon whatsapp-share" href="whatsapp://send?text=http://www.glassdoor.co.uk/Interview/Tesla-Interview-RVW44944117.htm" data-reviewid="44944117" data-action="share/whatsapp/share" data-label="whatsapp" rel="noopener noreferrer" target="_blank"><span class="offScreen">Share on WhatsApp</span></a></li>
					<li class="cell middle"><a class="social-share-icon email-share" href="mailto:?Subject=Tesla review on Glassdoor&amp;body=Read this review of Tesla on Glassdoor.  %22Mechanical Engineer Intern Interview%22&nbsp;http://www.glassdoor.co.uk/Interview/Tesla-Interview-RVW44944117.htm" data-label="email" data-reviewid="44944117" rel="noopener noreferrer" target="_blank"><span class="offScreen">Share via Email</span></a></li>
					<li class="cell middle"><a class="social-share-icon link-share" href="http://www.glassdoor.co.uk/Interview/Tesla-Interview-RVW44944117.htm" data-reviewid="44944117" data-label="link"><span class="offScreen">Copy link</span></a></li>
					<li class="cell linkCopySuccess"><span class="social-share-icon icon-check showDesk"></span><span>Link Copied!</span></li>
				</ul>
			</div>
		</div>
	</div>
	<div class="css-1dach6o d-flex align-items-center mt-std mt-md-0 justify-content-between justify-content-md-end">
		<button class="gd-ui-button mr-std css-glrvaa">Helpful</button>
		<div class="css-79elbk ewvknk0">
			<button class=" css-hhzi0d ewvknk1">
				<span class="SVGInline">
					<svg class="SVGInline-svg" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
						<path d="M9 2a9.3 9.3 0 00-4 1V2H4v20h1v-8c5.92-3.9 9.47 1.47 15 0V3c-4.17 1.1-7.21-1.09-11-1zm10 11.2c-1.38.22-2.65 0-5-.75l-.43-.13A13.8 13.8 0 009 11.54a7.45 7.45 0 00-4 1.26V4.2A7.71 7.71 0 019 3a16.39 16.39 0 014 .59h.14a14.42 14.42 0 005.86.64z" fill="currentColor" fill-rule="evenodd"></path>
					</svg>
				</span>
				<span class="d-none">Flag as Inappropriate</span>
			</button>
		</div>
	</div>
</div>
</div>
`

func TestParseInterview(t *testing.T) {
	want := Interview{
		ID:          "44944117",
		Date:        mustParseTime(t, "2021-04-02T00:00:00Z"),
		Title:       "Mechanical Engineer Intern Interview",
		Application: []string{"I interviewed at Tesla"},
		Process:     []string{"Highly flexible depending on team and directly interviewed by the team member, so could be just book technician questions or design scenarios. The number of times you get interviewed is also dependent on the team."},
		Questions:   []string{"Why do you want to work for Tesla?"},
	}

	root, err := html.Parse(strings.NewReader(fixture))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	got, err := parseInterview(root)
	if err != nil {
		t.Errorf("parseInterview() error = %v", err)
		return
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parseInterview() mismatch (-want +got):\n%s", diff)
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
