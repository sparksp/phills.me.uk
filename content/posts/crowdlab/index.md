+++
date = "2016-11-02T09:11:20Z"
title = "CrowdLab"
categories = ["programming", "project"]
aliases = [
  "/project/crowdlab"
]
toc.enable = false

[[resources]]
  name = "featured-image"
  src = "crowdlab.jpg"
  title = "Screenshot of CrowdLab website"
+++

Initially, I migrated the CrowdLab API from CodeIgniter (PHP) & MySQL to Ruby on Rails & PostgreSQL - re-engineering many of the features based on the company’s experiences of having run the system for a year.

One of the biggest improvements was enabling content additions and amends to be made by the client services team through a web interface. In the early versions of CrowdLab, such work was done by the development team.

After the migration was complete, I designed and implemented new features from the API / Database point of view, and considered the wider reaching effects of features.

I also managed the cloud servers for the staging and production environments, working toward automating as many processes as possible with Ansible and Jenkins.
