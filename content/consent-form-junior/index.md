+++
title = "Junior Consent Form"
description = "Parent and legal guardian consent form for participants under 18 years old taking part in climbing activities with Phill Sparks."
type = "page"
toc = false
meta = false
[sitemap]
  disable = true
aliases = ["consent-form/junior"]
+++

This consent form is for anyone under 18 years old and must be completed by the participant's parent or legal guardian. If the participant is over 18 years old please complete the [Adult Consent Form]({{< ref "/consent-form" >}}).

{{< form name="consent-form-junior" action="/consent-form-success/" >}}

  {{< form/fieldset legend="Participant Information" >}}
    <p>Please complete all required fields prior to attending the session or course.</p>

    {{< form/text name="first-name" label="First Name" required="yes" autocomplete="section-participant given-name" autofocus="yes" >}}
    {{< form/text name="last-name" label="Last Name" required="yes" autocomplete="section-participant family-name" >}}
    {{< form/email name="email" label="E-mail Address" required="yes" autocomplete="section-participant email" >}}
    {{< form/tel name="phone-number" label="Phone Number" required="yes" autocomplete="section-participant tel" >}}
    {{< form/date name="date-of-birth" label="Date of Birth" placeholder="dd/mm/yyyy" required="yes" autocomplete="section-participant bday" >}}
    {{< form/text name="post-code" label="Postcode" required="yes" autocomplete="section-participant postal-code" >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Medical Details" >}}
    <p>Please detail below any physical, mental, or medical conditions that may affect the participant's ability to participate in the activities being provided, or state <strong>"none"</strong> if there are no issues. <em>(Please contact me beforehand if you are unsure).</em></p>

    {{< form/textarea name="medical-details" label="Medical Details" required="yes" autocomplete="off" >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Emergency Contact" >}}
    <p>Please provide the name and phone number of someone who can be contacted in the event of an emergency if you are unable to be reached directly.</p>

    {{< form/text name="emergency-name" label="Emergency Contact Name" required="yes" autocomplete="section-emergency name" >}}
    {{< form/tel name="emergency-tel" label="Emergency Contact Number" required="yes" autocomplete="section-emergency tel" >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Participation Statement" >}}
    <p>Climbing can be mentally and physically demanding, with inherent risks and hazards. Activities are run and supervised by appropriately qualified and experienced instructors. All necessary safety equipment is provided, regularly checked, and maintained. Minor injuries (e.g., cuts and bruises) are possible, and more serious injuries cannot be entirely ruled out. Instructors strive to maintain a safe environment—please ensure the participant follows instructions, asks questions when needed, and brings suitable outdoor clothing, food, and drink. Alcohol or drug intoxication is strictly prohibited.</p>

    <p>The British Mountaineering Council (BMC) recognises that climbing and mountaineering are activities with a danger of personal injury or death. Participants in these activities should be aware of and accept these risks and be responsible for their own actions.</p>

    {{< form/checkbox name="participation-statement" required="yes" >}}
      By ticking this box, you agree that the participant details above are correct and that you understand and accept this Participation Statement.
    {{< /form/checkbox >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Electronic Signature & Data Protection" >}}
    <p>To comply with UK GDPR regulations, personal data is collected and stored securely by Phill Sparks as Data Controller. Data is retained for 3 years from the participant's most recent session to satisfy insurance requirements. You have the right to request access to stored data or request erasure where applicable under law. Full details are available in the <a href="{{< ref "/policies/privacy" >}}">Privacy Policy</a>.</p>

    {{< form/checkbox name="electronic-waiver" required="yes" >}}
      I consent to using an electronic signature in lieu of a paper signature on behalf of the participant.
    {{< /form/checkbox >}}

    {{< form/checkbox name="data-protection-policy" required="yes" >}}
      I agree to the Data Protection Policy and Privacy Policy.
    {{< /form/checkbox >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Parental Declaration" >}}
    <p>I am the parent or legal guardian of the participant named above and understand the terms of the booking. I have read and fully agree with the terms, policies, and safety instructions provided.</p>

    {{< form/text name="declaration-full-name" label="Print Full Name" required="yes" autocomplete="section-declaration name" >}}
    {{< form/text name="declaration-relationship" label="Relationship to Participant" placeholder="Parent or Legal Guardian" required="yes" >}}
    {{< form/date name="declaration-date" label="Today's Date" placeholder="dd/mm/yyyy" required="yes" >}}
  {{< /form/fieldset >}}

  {{< form/recaptcha >}}
  {{< form/submit label="Submit Consent Form" >}}

{{< /form >}}